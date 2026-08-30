package list

import (
	"context"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/types"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func ListQuests(ctx context.Context, req ListQuestsRequest) *ListQuestsResponse {
	config := common.GetConfig()

	res := &ListQuestsResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_ListQuests,
		},
		StatusCode: ListQuestsCodes_Unknown,
	}

	cursor, err := collections.GetQuestsCollection().Find(ctx, bson.M{})
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.StatusCode = ListQuestsCodes_FailedToListQuests
		res.Message = "Failed to list quests"
		return res
	}

	quests := []types.Quest{}
	if err := cursor.All(ctx, &quests); err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.StatusCode = ListQuestsCodes_FailedToListQuests
		res.Message = "Failed to list quests"
		return res
	}

	res.HttpCode = common.HttpCodes_OK
	res.StatusCode = ListQuestsCodes_Success
	res.Message = "Quests listed successfully"
	res.Quests = quests
	return res
}
