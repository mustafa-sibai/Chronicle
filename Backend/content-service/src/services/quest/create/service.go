package create

import (
	"context"
	"time"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/types"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateQuest(ctx context.Context, req CreateQuestRequest) *CreateQuestResponse {
	config := common.GetConfig()

	res := &CreateQuestResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_CreateQuest,
		},
		StatusCode: CreateQuestCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = CreateQuestCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	itemTemplateIDs, err := convertMongoDBStringIDsToObjectIDs(req.Body.Reward.ItemTemplateIDs)
	if err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = CreateQuestCodes_InvalidInput
		res.Message = "reward.itemTemplateIds contains an invalid id"
		return res
	}

	choiceItemTemplateIDs, err := convertMongoDBStringIDsToObjectIDs(req.Body.Reward.ChoiceItemTemplateIDs)
	if err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = CreateQuestCodes_InvalidInput
		res.Message = "reward.choiceItemTemplateIds contains an invalid id"
		return res
	}

	allItemTemplateIDs := append(append([]bson.ObjectID{}, itemTemplateIDs...), choiceItemTemplateIDs...)
	if len(allItemTemplateIDs) > 0 {
		count, err := collections.GetItemTemplatesCollection().CountDocuments(ctx, bson.M{"_id": bson.M{"$in": allItemTemplateIDs}})
		if err != nil {
			res.HttpCode = common.HttpCodes_InternalServerError
			res.Message = "Failed to create quest"
			return res
		}
		if int(count) != len(uniqueObjectIDs(allItemTemplateIDs)) {
			res.HttpCode = common.HttpCodes_BadRequest
			res.StatusCode = CreateQuestCodes_RewardItemTemplateNotFound
			res.Message = "One or more reward item templates were not found"
			return res
		}
	}

	quest := types.Quest{
		Title:       req.Body.Title,
		Description: req.Body.Description,
		Objective: types.QuestObjective{
			Description:     req.Body.Objective.Description,
			CompletionCount: req.Body.Objective.CompletionCount,
		},
		Reward: types.QuestReward{
			ItemTemplateIDs:       itemTemplateIDs,
			ChoiceItemTemplateIDs: choiceItemTemplateIDs,
			Experience:            req.Body.Reward.Experience,
		},
		CreatedAt: time.Now().UTC(),
	}

	result, err := collections.GetQuestsCollection().InsertOne(ctx, quest)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.StatusCode = CreateQuestCodes_FailedToCreateQuest
		res.Message = "Failed to create quest"
		return res
	}

	res.HttpCode = common.HttpCodes_Created
	res.StatusCode = CreateQuestCodes_QuestCreatedSuccessfully
	res.Message = "Quest created successfully"
	res.QuestID = result.InsertedID.(bson.ObjectID).Hex()
	return res
}

func convertMongoDBStringIDsToObjectIDs(hexIDs []string) ([]bson.ObjectID, error) {
	ids := make([]bson.ObjectID, 0, len(hexIDs))
	for _, hexID := range hexIDs {
		id, err := bson.ObjectIDFromHex(hexID)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func uniqueObjectIDs(ids []bson.ObjectID) []bson.ObjectID {
	seen := make(map[bson.ObjectID]struct{}, len(ids))
	unique := make([]bson.ObjectID, 0, len(ids))
	for _, id := range ids {
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		unique = append(unique, id)
	}
	return unique
}
