package delete

import (
	"context"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func DeleteItemTemplate(ctx context.Context, req DeleteItemTemplateRequest) *DeleteItemTemplateResponse {
	config := common.GetConfig()

	res := &DeleteItemTemplateResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_DeleteItemTemplate,
		},
		StatusCode: DeleteItemTemplateCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = DeleteItemTemplateCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	itemTemplateObjectID, err := bson.ObjectIDFromHex(req.Body.ItemTemplateID)
	if err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = DeleteItemTemplateCodes_InvalidInput
		res.Message = "itemTemplateId is invalid"
		return res
	}

	itemTemplatesCollection := collections.GetItemTemplatesCollection()

	result, err := itemTemplatesCollection.DeleteOne(ctx, bson.M{"_id": itemTemplateObjectID})
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.StatusCode = DeleteItemTemplateCodes_FailedToDeleteItemTemplate
		res.Message = "Failed to delete item template"
		return res
	}
	if result.DeletedCount == 0 {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = DeleteItemTemplateCodes_ItemTemplateNotFound
		res.Message = "Item template not found"
		return res
	}

	res.HttpCode = common.HttpCodes_OK
	res.StatusCode = DeleteItemTemplateCodes_ItemTemplateDeletedSuccessfully
	res.Message = "Item template deleted successfully"
	return res
}
