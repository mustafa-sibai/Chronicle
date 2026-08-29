package get

import (
	"context"
	"errors"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/game"
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

	templateObjectID, err := bson.ObjectIDFromHex(req.Body.TemplateID)
	if err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = GetItemTemplateCodes_InvalidInput
		res.Message = "templateId is invalid"
		return res
	}

	itemTemplatesCollection := collections.GetItemTemplatesCollection()

	raw, err := itemTemplatesCollection.FindOne(ctx, bson.M{"_id": templateObjectID}).Raw()
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

	var probe struct {
		ItemType game.ItemType `bson:"itemType"`
	}
	if err := bson.Unmarshal(raw, &probe); err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.StatusCode = GetItemTemplateCodes_FailedToGetItemTemplate
		res.Message = "Failed to get item template"
		return res
	}

	var itemTemplate any
	switch probe.ItemType {
	case game.ItemType_Container:
		var containerItemTemplate types.ContainerItemTemplate
		err = bson.Unmarshal(raw, &containerItemTemplate)
		itemTemplate = containerItemTemplate
	case game.ItemType_Weapon, game.ItemType_Armor:
		var equipmentItemTemplate types.EquipmentItemTemplate
		err = bson.Unmarshal(raw, &equipmentItemTemplate)
		itemTemplate = equipmentItemTemplate
	default:
		var plainItemTemplate types.ItemTemplate
		err = bson.Unmarshal(raw, &plainItemTemplate)
		itemTemplate = plainItemTemplate
	}
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
