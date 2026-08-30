package get

import (
	"context"
	"errors"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/types"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetQuest(ctx context.Context, req GetQuestRequest) *GetQuestResponse {
	config := common.GetConfig()

	res := &GetQuestResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_GetQuest,
		},
		StatusCode: GetQuestCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = GetQuestCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	questObjectID, err := bson.ObjectIDFromHex(req.Body.QuestID)
	if err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = GetQuestCodes_InvalidInput
		res.Message = "questId is invalid"
		return res
	}

	var quest types.Quest
	err = collections.GetQuestsCollection().FindOne(ctx, bson.M{"_id": questObjectID}).Decode(&quest)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			res.HttpCode = common.HttpCodes_BadRequest
			res.StatusCode = GetQuestCodes_QuestNotFound
			res.Message = "Quest not found"
			return res
		}
		res.HttpCode = common.HttpCodes_InternalServerError
		res.StatusCode = GetQuestCodes_FailedToGetQuest
		res.Message = "Failed to get quest"
		return res
	}

	res.HttpCode = common.HttpCodes_OK
	res.StatusCode = GetQuestCodes_Success
	res.Message = "Quest retrieved successfully"
	res.Quest = &quest
	return res
}
