package create

import (
	"context"
	"errors"
	"time"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/game"
	"github.com/mustafa-sibai/chronicle/backend-lib/session"
	"github.com/mustafa-sibai/chronicle/backend-lib/types"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateCharacter(ctx context.Context, req CreateCharacterRequest) *CreateCharacterResponse {
	config := common.GetConfig()

	res := &CreateCharacterResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_CreateCharacter,
		},
		StatusCode: CreateCharacterCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = CreateCharacterCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	if !req.Body.Race.IsClassAllowedInRace(req.Body.Class) {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = CreateCharacterCodes_InvalidRaceClassCombination
		res.Message = "class is not available for the selected race"
		return res
	}

	accountID, ok, err := session.ResolveAccountID(ctx, req.Head.SessionID)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to create character"
		return res
	}
	if !ok {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = CreateCharacterCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	accountObjectID, err := bson.ObjectIDFromHex(accountID)
	if err != nil {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = CreateCharacterCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	charactersCollection := collections.GetCharactersCollection()

	count, err := charactersCollection.CountDocuments(ctx, bson.M{"accountId": accountObjectID})
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to create character"
		return res
	}
	if count >= game.MaxCharactersPerAccount {
		res.HttpCode = common.HttpCodes_Conflict
		res.StatusCode = CreateCharacterCodes_CharacterLimitReached
		res.Message = "Character limit reached"
		return res
	}

	err = charactersCollection.FindOne(ctx, bson.M{"name": req.Body.Name}).Err()
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to create character"
		return res
	}
	if err == nil {
		res.HttpCode = common.HttpCodes_Conflict
		res.StatusCode = CreateCharacterCodes_NameTaken
		res.Message = "Character name already taken"
		return res
	}

	character := types.Character{
		AccountID:  accountObjectID,
		Name:       req.Body.Name,
		Race:       req.Body.Race,
		Class:      req.Body.Class,
		Gender:     req.Body.Gender,
		Appearance: req.Body.Appearance,
		CreatedAt:  time.Now().UTC(),
	}

	result, err := charactersCollection.InsertOne(ctx, character)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to create character"
		return res
	}

	res.HttpCode = common.HttpCodes_Created
	res.StatusCode = CreateCharacterCodes_CharacterCreatedSuccessfully
	res.Message = "Character created successfully"
	res.CharacterID = result.InsertedID.(bson.ObjectID).Hex()
	return res
}
