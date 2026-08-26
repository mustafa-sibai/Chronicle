package delete

import (
	"context"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/session"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func DeleteCharacter(ctx context.Context, req DeleteCharacterRequest) *DeleteCharacterResponse {
	config := common.GetConfig()

	res := &DeleteCharacterResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_DeleteCharacter,
		},
		StatusCode: DeleteCharacterCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = DeleteCharacterCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	accountID, ok, err := session.ResolveAccountID(ctx, req.Head.SessionID)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to delete character"
		return res
	}
	if !ok {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = DeleteCharacterCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	accountObjectID, err := bson.ObjectIDFromHex(accountID)
	if err != nil {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = DeleteCharacterCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	characterObjectID, err := bson.ObjectIDFromHex(req.Body.CharacterID)
	if err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = DeleteCharacterCodes_InvalidInput
		res.Message = "characterId is invalid"
		return res
	}

	charactersCollection := collections.GetCharactersCollection()

	result, err := charactersCollection.DeleteOne(ctx, bson.M{"_id": characterObjectID, "accountId": accountObjectID})
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to delete character"
		return res
	}
	if result.DeletedCount == 0 {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = DeleteCharacterCodes_CharacterNotFound
		res.Message = "Character not found"
		return res
	}

	res.HttpCode = common.HttpCodes_OK
	res.StatusCode = DeleteCharacterCodes_CharacterDeletedSuccessfully
	res.Message = "Character deleted successfully"
	return res
}
