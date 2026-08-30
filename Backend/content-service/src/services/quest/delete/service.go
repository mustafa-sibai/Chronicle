package delete

import (
	"context"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func DeleteQuest(ctx context.Context, req DeleteQuestRequest) *DeleteQuestResponse {
	config := common.GetConfig()

	res := &DeleteQuestResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_DeleteQuest,
		},
		StatusCode: DeleteQuestCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = DeleteQuestCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	questObjectID, err := bson.ObjectIDFromHex(req.Body.QuestID)
	if err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = DeleteQuestCodes_InvalidInput
		res.Message = "questId is invalid"
		return res
	}

	result, err := collections.GetQuestsCollection().DeleteOne(ctx, bson.M{"_id": questObjectID})
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.StatusCode = DeleteQuestCodes_FailedToDeleteQuest
		res.Message = "Failed to delete quest"
		return res
	}
	if result.DeletedCount == 0 {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = DeleteQuestCodes_QuestNotFound
		res.Message = "Quest not found"
		return res
	}

	res.HttpCode = common.HttpCodes_OK
	res.StatusCode = DeleteQuestCodes_QuestDeletedSuccessfully
	res.Message = "Quest deleted successfully"
	return res
}
