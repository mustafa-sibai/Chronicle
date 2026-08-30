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

func GetItemTemplate(ctx context.Context, req GetItemTemplateRequest) *GetItemTemplateResponse {
	config := common.GetConfig()

	res := &GetItemTemplateResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_GetItemTemplate,
		},
		StatusCode: GetItemTemplateCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = GetItemTemplateCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	itemTemplateObjectID, err := bson.ObjectIDFromHex(req.Body.ItemTemplateID)
	if err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = GetItemTemplateCodes_InvalidInput
		res.Message = "itemTemplateId is invalid"
		return res
	}

	raw, err := collections.GetItemTemplatesCollection().FindOne(ctx, bson.M{"_id": itemTemplateObjectID}).Raw()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			res.HttpCode = common.HttpCodes_BadRequest
			res.StatusCode = GetItemTemplateCodes_ItemTemplateNotFound
			res.Message = "Item template not found"
			return res
		}
		res.HttpCode = common.HttpCodes_InternalServerError
		res.StatusCode = GetItemTemplateCodes_FailedToGetItemTemplate
		res.Message = "Failed to get item template"
		return res
	}

	itemTemplate, err := types.DecodeItemTemplate(raw)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.StatusCode = GetItemTemplateCodes_FailedToGetItemTemplate
		res.Message = "Failed to get item template"
		return res
	}

	res.HttpCode = common.HttpCodes_OK
	res.StatusCode = GetItemTemplateCodes_Success
	res.Message = "Item template retrieved successfully"
	res.ItemTemplate = itemTemplate
	return res
}
