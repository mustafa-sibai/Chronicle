package list

import (
	"context"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/session"
	"github.com/mustafa-sibai/chronicle/backend-lib/types"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func ListCharacters(ctx context.Context, req ListCharactersRequest) *ListCharactersResponse {
	config := common.GetConfig()

	res := &ListCharactersResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_ListCharacters,
		},
		StatusCode: ListCharactersCodes_Unknown,
	}

	accountID, ok, err := session.ResolveAccountID(ctx, req.Head.SessionID)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to list characters"
		return res
	}
	if !ok {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = ListCharactersCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	accountObjectID, err := bson.ObjectIDFromHex(accountID)
	if err != nil {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = ListCharactersCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	charactersCollection := collections.GetCharactersCollection()

	cursor, err := charactersCollection.Find(ctx, bson.M{"accountId": accountObjectID})
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to list characters"
		return res
	}

	characters := []types.Character{}
	if err := cursor.All(ctx, &characters); err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to list characters"
		return res
	}

	res.HttpCode = common.HttpCodes_OK
	res.StatusCode = ListCharactersCodes_Success
	res.Message = "Characters listed successfully"
	res.Characters = characters
	return res
}
