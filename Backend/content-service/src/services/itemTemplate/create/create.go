package create

import (
	"context"
	"time"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/game"
	"github.com/mustafa-sibai/chronicle/backend-lib/types"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateItemTemplate(ctx context.Context, req CreateItemTemplateRequest) *CreateItemTemplateResponse {
	config := common.GetConfig()

	res := &CreateItemTemplateResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_CreateItemTemplate,
		},
		StatusCode: CreateItemTemplateCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = CreateItemTemplateCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	itemTemplate := types.ItemTemplate{
		Name:      req.Body.Name,
		ItemType:  req.Body.ItemType,
		MaxStacks: req.Body.MaxStacks,
		CreatedAt: time.Now().UTC(),
	}

	var itemTemplateDoc any = itemTemplate
	switch req.Body.ItemType {
	case game.ItemType_Container:
		itemTemplateDoc = types.ContainerItemTemplate{
			ItemTemplate: itemTemplate,
			Capacity:     req.Body.Capacity,
		}
	case game.ItemType_Weapon, game.ItemType_Armor:
		itemTemplateDoc = types.EquipmentItemTemplate{
			ItemTemplate: itemTemplate,
			Stats:        *req.Body.Stats,
		}
	}

	itemTemplatesCollection := collections.GetItemTemplatesCollection()

	result, err := itemTemplatesCollection.InsertOne(ctx, itemTemplateDoc)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.StatusCode = CreateItemTemplateCodes_FailedToCreateItemTemplate
		res.Message = "Failed to create item template"
		return res
	}

	res.HttpCode = common.HttpCodes_Created
	res.StatusCode = CreateItemTemplateCodes_ItemTemplateCreatedSuccessfully
	res.Message = "Item template created successfully"
	res.TemplateID = result.InsertedID.(bson.ObjectID).Hex()
	return res
}
