package complete

import (
	"context"
	"errors"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/database/mongodb"
	"github.com/mustafa-sibai/chronicle/backend-lib/session"
	"github.com/mustafa-sibai/chronicle/backend-lib/types"
	"github.com/mustafa-sibai/chronicle/character-service/services/inventory/add"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CompleteQuest(ctx context.Context, req CompleteQuestRequest) *CompleteQuestResponse {
	config := common.GetConfig()

	res := &CompleteQuestResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_CompleteQuest,
		},
		StatusCode: CompleteQuestCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = CompleteQuestCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	accountID, ok, err := session.ResolveAccountID(ctx, req.Head.SessionID)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to complete quest"
		return res
	}
	if !ok {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = CompleteQuestCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	accountObjectID, err := bson.ObjectIDFromHex(accountID)
	if err != nil {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = CompleteQuestCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	characterID, ok, err := session.ResolveCharacterID(ctx, req.Head.SessionID)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to complete quest"
		return res
	}
	if !ok {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = CompleteQuestCodes_NoCharacterActive
		res.Message = "No character is currently active on this session"
		return res
	}

	characterObjectID, err := bson.ObjectIDFromHex(characterID)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to complete quest"
		return res
	}

	questObjectID, err := bson.ObjectIDFromHex(req.Body.QuestID)
	if err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = CompleteQuestCodes_InvalidInput
		res.Message = "questId is invalid"
		return res
	}

	charactersCollection := collections.GetCharactersCollection()

	var character types.Character
	err = charactersCollection.FindOne(ctx, bson.M{"_id": characterObjectID, "accountId": accountObjectID}).Decode(&character)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			res.HttpCode = common.HttpCodes_BadRequest
			res.StatusCode = CompleteQuestCodes_CharacterNotFound
			res.Message = "Character not found"
			return res
		}
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to complete quest"
		return res
	}

	entryIndex := -1
	for i, activeQuest := range character.ActiveQuests {
		if activeQuest.QuestID == questObjectID {
			entryIndex = i
			break
		}
	}
	if entryIndex == -1 {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = CompleteQuestCodes_QuestNotActive
		res.Message = "Quest is not active for this character"
		return res
	}

	var quest types.Quest
	err = collections.GetQuestsCollection().FindOne(ctx, bson.M{"_id": questObjectID}).Decode(&quest)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to complete quest"
		return res
	}

	if character.ActiveQuests[entryIndex].Progress < quest.Objective.CompletionCount {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = CompleteQuestCodes_ObjectiveNotMet
		res.Message = "Quest objective not yet met"
		return res
	}

	rewardItemTemplateIDs := append([]bson.ObjectID{}, quest.Reward.ItemTemplateIDs...)

	if len(quest.Reward.ChoiceItemTemplateIDs) > 0 {
		if req.Body.ChosenRewardItemTemplateID == "" {
			res.HttpCode = common.HttpCodes_BadRequest
			res.StatusCode = CompleteQuestCodes_RewardChoiceRequired
			res.Message = "chosenRewardItemTemplateId is required for this quest"
			return res
		}
		chosenObjectID, err := bson.ObjectIDFromHex(req.Body.ChosenRewardItemTemplateID)
		if err != nil {
			res.HttpCode = common.HttpCodes_BadRequest
			res.StatusCode = CompleteQuestCodes_InvalidInput
			res.Message = "chosenRewardItemTemplateId is invalid"
			return res
		}
		offered := false
		for _, choiceID := range quest.Reward.ChoiceItemTemplateIDs {
			if choiceID == chosenObjectID {
				offered = true
				break
			}
		}
		if !offered {
			res.HttpCode = common.HttpCodes_BadRequest
			res.StatusCode = CompleteQuestCodes_InvalidRewardChoice
			res.Message = "chosenRewardItemTemplateId is not one of the quest's reward choices"
			return res
		}
		rewardItemTemplateIDs = append(rewardItemTemplateIDs, chosenObjectID)
	} else if req.Body.ChosenRewardItemTemplateID != "" {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = CompleteQuestCodes_InvalidInput
		res.Message = "this quest has no reward choice to make"
		return res
	}

	var inventory types.Inventory
	err = collections.GetInventoriesCollection().FindOne(ctx, bson.M{"characterId": characterObjectID}).Decode(&inventory)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to complete quest"
		return res
	}

	// Place every reward item on an in-memory copy of the inventory before
	// persisting anything, so a failure partway through commits nothing.
	for _, itemTemplateID := range rewardItemTemplateIDs {
		var rewardItemTemplate types.ItemTemplate
		err = collections.GetItemTemplatesCollection().FindOne(ctx, bson.M{"_id": itemTemplateID}).Decode(&rewardItemTemplate)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				res.HttpCode = common.HttpCodes_InternalServerError
				res.StatusCode = CompleteQuestCodes_RewardItemTemplateNotFound
				res.Message = "A reward item template no longer exists"
				return res
			}
			res.HttpCode = common.HttpCodes_InternalServerError
			res.Message = "Failed to complete quest"
			return res
		}

		remaining, err := add.PlaceItem(ctx, &inventory, itemTemplateID, rewardItemTemplate.ItemType, rewardItemTemplate.MaxStacks, 1)
		if err != nil {
			res.HttpCode = common.HttpCodes_InternalServerError
			res.Message = "Failed to complete quest"
			return res
		}
		if remaining > 0 {
			res.HttpCode = common.HttpCodes_Conflict
			res.StatusCode = CompleteQuestCodes_InventoryFull
			res.Message = "Inventory full"
			return res
		}
	}

	mongoSession, err := mongodb.GetMongoDBClient().StartSession()
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to complete quest"
		return res
	}
	defer mongoSession.EndSession(ctx)

	_, err = mongoSession.WithTransaction(ctx, func(sessCtx context.Context) (any, error) {
		if _, err := collections.GetInventoriesCollection().ReplaceOne(sessCtx, bson.M{"_id": inventory.ID}, inventory); err != nil {
			return nil, err
		}

		update := bson.M{
			"$pull": bson.M{"activeQuests": bson.M{"questId": questObjectID}},
			"$push": bson.M{"completedQuests": questObjectID},
			"$inc":  bson.M{"experience": quest.Reward.Experience},
		}
		if _, err := charactersCollection.UpdateByID(sessCtx, characterObjectID, update); err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to complete quest"
		return res
	}

	res.HttpCode = common.HttpCodes_OK
	res.StatusCode = CompleteQuestCodes_QuestCompletedSuccessfully
	res.Message = "Quest completed successfully"
	return res
}
