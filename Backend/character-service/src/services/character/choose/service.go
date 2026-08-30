package choose

import (
	"context"
	"errors"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/database/valkeydb"
	"github.com/mustafa-sibai/chronicle/backend-lib/session"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func ChooseCharacter(ctx context.Context, req ChooseCharacterRequest) *ChooseCharacterResponse {
	config := common.GetConfig()

	res := &ChooseCharacterResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_ChooseCharacter,
		},
		StatusCode: ChooseCharacterCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = ChooseCharacterCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	accountID, ok, err := session.ResolveAccountID(ctx, req.Head.SessionID)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to choose character"
		return res
	}
	if !ok {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = ChooseCharacterCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	accountObjectID, err := bson.ObjectIDFromHex(accountID)
	if err != nil {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = ChooseCharacterCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	characterObjectID, err := bson.ObjectIDFromHex(req.Body.CharacterID)
	if err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = ChooseCharacterCodes_InvalidInput
		res.Message = "characterId is invalid"
		return res
	}

	err = collections.GetCharactersCollection().FindOne(ctx, bson.M{"_id": characterObjectID, "accountId": accountObjectID}).Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			res.HttpCode = common.HttpCodes_BadRequest
			res.StatusCode = ChooseCharacterCodes_CharacterNotFound
			res.Message = "Character not found"
			return res
		}
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to choose character"
		return res
	}

	valkeyClient := valkeydb.GetValkeyClient()
	setCmd := valkeyClient.B().Hsetex().Key(session.SessionKey(req.Head.SessionID)).
		Ex(int64(session.SessionTTL.Seconds())).
		Fields().Numfields(1).
		FieldValue().FieldValue("characterId", req.Body.CharacterID).
		Build()
	if err := valkeyClient.Do(ctx, setCmd).Error(); err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.StatusCode = ChooseCharacterCodes_FailedToChooseCharacter
		res.Message = "Failed to choose character"
		return res
	}

	res.HttpCode = common.HttpCodes_OK
	res.StatusCode = ChooseCharacterCodes_CharacterChosenSuccessfully
	res.Message = "Character chosen successfully"
	return res
}
