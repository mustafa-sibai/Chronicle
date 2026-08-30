package get

import (
	"context"
	"errors"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/session"
	"github.com/mustafa-sibai/chronicle/backend-lib/types"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetInventory(ctx context.Context, req GetInventoryRequest) *GetInventoryResponse {
	config := common.GetConfig()

	res := &GetInventoryResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_GetInventory,
		},
		StatusCode: GetInventoryCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = GetInventoryCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	accountID, ok, err := session.ResolveAccountID(ctx, req.Head.SessionID)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to get inventory"
		return res
	}
	if !ok {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = GetInventoryCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	accountObjectID, err := bson.ObjectIDFromHex(accountID)
	if err != nil {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = GetInventoryCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	characterID, ok, err := session.ResolveCharacterID(ctx, req.Head.SessionID)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to get inventory"
		return res
	}
	if !ok {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = GetInventoryCodes_NoCharacterActive
		res.Message = "No character is currently active on this session"
		return res
	}

	characterObjectID, err := bson.ObjectIDFromHex(characterID)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to get inventory"
		return res
	}

	err = collections.GetCharactersCollection().FindOne(ctx, bson.M{"_id": characterObjectID, "accountId": accountObjectID}).Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			res.HttpCode = common.HttpCodes_BadRequest
			res.StatusCode = GetInventoryCodes_CharacterNotFound
			res.Message = "Character not found"
			return res
		}
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to get inventory"
		return res
	}

	var inventory types.Inventory
	err = collections.GetInventoriesCollection().FindOne(ctx, bson.M{"characterId": characterObjectID}).Decode(&inventory)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to get inventory"
		return res
	}

	res.HttpCode = common.HttpCodes_OK
	res.StatusCode = GetInventoryCodes_Success
	res.Message = "Inventory retrieved successfully"
	res.Inventory = &inventory
	return res
}
