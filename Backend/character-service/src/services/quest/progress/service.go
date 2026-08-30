package progress

import (
	"context"
	"errors"
	"strconv"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/session"
	"github.com/mustafa-sibai/chronicle/backend-lib/types"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func UpdateQuestProgress(ctx context.Context, req UpdateQuestProgressRequest) *UpdateQuestProgressResponse {
	config := common.GetConfig()

	res := &UpdateQuestProgressResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_UpdateQuestProgress,
		},
		StatusCode: UpdateQuestProgressCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = UpdateQuestProgressCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	accountID, ok, err := session.ResolveAccountID(ctx, req.Head.SessionID)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to update quest progress"
		return res
	}
	if !ok {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = UpdateQuestProgressCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	accountObjectID, err := bson.ObjectIDFromHex(accountID)
	if err != nil {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = UpdateQuestProgressCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	characterID, ok, err := session.ResolveCharacterID(ctx, req.Head.SessionID)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to update quest progress"
		return res
	}
	if !ok {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = UpdateQuestProgressCodes_NoCharacterActive
		res.Message = "No character is currently active on this session"
		return res
	}

	characterObjectID, err := bson.ObjectIDFromHex(characterID)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to update quest progress"
		return res
	}

	questObjectID, err := bson.ObjectIDFromHex(req.Body.QuestID)
	if err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = UpdateQuestProgressCodes_InvalidInput
		res.Message = "questId is invalid"
		return res
	}

	charactersCollection := collections.GetCharactersCollection()

	var character types.Character
	err = charactersCollection.FindOne(ctx, bson.M{"_id": characterObjectID, "accountId": accountObjectID}).Decode(&character)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			res.HttpCode = common.HttpCodes_BadRequest
			res.StatusCode = UpdateQuestProgressCodes_CharacterNotFound
			res.Message = "Character not found"
			return res
		}
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to update quest progress"
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
		res.StatusCode = UpdateQuestProgressCodes_QuestNotActive
		res.Message = "Quest is not active for this character"
		return res
	}

	var quest types.Quest
	err = collections.GetQuestsCollection().FindOne(ctx, bson.M{"_id": questObjectID}).Decode(&quest)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to update quest progress"
		return res
	}

	newProgress := min(character.ActiveQuests[entryIndex].Progress+req.Body.Amount, quest.Objective.CompletionCount)

	field := "activeQuests." + strconv.Itoa(entryIndex) + ".progress"
	update := bson.M{"$set": bson.M{field: newProgress}}
	if _, err := charactersCollection.UpdateByID(ctx, characterObjectID, update); err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to update quest progress"
		return res
	}

	res.HttpCode = common.HttpCodes_OK
	res.StatusCode = UpdateQuestProgressCodes_ProgressUpdatedSuccessfully
	res.Message = "Quest progress updated successfully"
	res.Progress = newProgress
	return res
}
