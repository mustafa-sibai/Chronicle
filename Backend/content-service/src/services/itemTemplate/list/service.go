package list

import (
	"context"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/types"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func ListItemTemplates(ctx context.Context, req ListItemTemplatesRequest) *ListItemTemplatesResponse {
	config := common.GetConfig()

	res := &ListItemTemplatesResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_ListItemTemplates,
		},
		StatusCode: ListItemTemplatesCodes_Unknown,
	}

	cursor, err := collections.GetItemTemplatesCollection().Find(ctx, bson.M{})
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.StatusCode = ListItemTemplatesCodes_FailedToListItemTemplates
		res.Message = "Failed to list item templates"
		return res
	}
	defer cursor.Close(ctx)

	itemTemplates := []any{}
	for cursor.Next(ctx) {
		itemTemplate, err := types.DecodeItemTemplate(cursor.Current)
		if err != nil {
			res.HttpCode = common.HttpCodes_InternalServerError
			res.StatusCode = ListItemTemplatesCodes_FailedToListItemTemplates
			res.Message = "Failed to list item templates"
			return res
		}
		itemTemplates = append(itemTemplates, itemTemplate)
	}
	if err := cursor.Err(); err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.StatusCode = ListItemTemplatesCodes_FailedToListItemTemplates
		res.Message = "Failed to list item templates"
		return res
	}

	res.HttpCode = common.HttpCodes_OK
	res.StatusCode = ListItemTemplatesCodes_Success
	res.Message = "Item templates listed successfully"
	res.ItemTemplates = itemTemplates
	return res
}
