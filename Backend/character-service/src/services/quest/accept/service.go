package accept

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

func AcceptQuest(ctx context.Context, req AcceptQuestRequest) *AcceptQuestResponse {
	config := common.GetConfig()

	res := &AcceptQuestResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_AcceptQuest,
		},
		StatusCode: AcceptQuestCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = AcceptQuestCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	accountID, ok, err := session.ResolveAccountID(ctx, req.Head.SessionID)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to accept quest"
		return res
	}
	if !ok {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = AcceptQuestCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	accountObjectID, err := bson.ObjectIDFromHex(accountID)
	if err != nil {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = AcceptQuestCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	characterID, ok, err := session.ResolveCharacterID(ctx, req.Head.SessionID)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to accept quest"
		return res
	}
	if !ok {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = AcceptQuestCodes_NoCharacterActive
		res.Message = "No character is currently active on this session"
		return res
	}

	characterObjectID, err := bson.ObjectIDFromHex(characterID)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to accept quest"
		return res
	}

	questObjectID, err := bson.ObjectIDFromHex(req.Body.QuestID)
	if err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = AcceptQuestCodes_InvalidInput
		res.Message = "questId is invalid"
		return res
	}

	err = collections.GetQuestsCollection().FindOne(ctx, bson.M{"_id": questObjectID}).Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			res.HttpCode = common.HttpCodes_BadRequest
			res.StatusCode = AcceptQuestCodes_QuestNotFound
			res.Message = "Quest not found"
			return res
		}
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to accept quest"
		return res
	}

	charactersCollection := collections.GetCharactersCollection()

	var character types.Character
	err = charactersCollection.FindOne(ctx, bson.M{"_id": characterObjectID, "accountId": accountObjectID}).Decode(&character)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			res.HttpCode = common.HttpCodes_BadRequest
			res.StatusCode = AcceptQuestCodes_CharacterNotFound
			res.Message = "Character not found"
			return res
		}
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to accept quest"
		return res
	}

	for _, completedQuestID := range character.CompletedQuests {
		if completedQuestID == questObjectID {
			res.HttpCode = common.HttpCodes_Conflict
			res.StatusCode = AcceptQuestCodes_QuestAlreadyCompleted
			res.Message = "Quest already completed"
			return res
		}
	}
	for _, activeQuest := range character.ActiveQuests {
		if activeQuest.QuestID == questObjectID {
			res.HttpCode = common.HttpCodes_Conflict
			res.StatusCode = AcceptQuestCodes_QuestAlreadyActive
			res.Message = "Quest already active"
			return res
		}
	}

	entry := types.QuestLogEntry{QuestID: questObjectID, Progress: 0}
	update := bson.M{"$push": bson.M{"activeQuests": entry}}
	if _, err := charactersCollection.UpdateByID(ctx, characterObjectID, update); err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to accept quest"
		return res
	}

	res.HttpCode = common.HttpCodes_OK
	res.StatusCode = AcceptQuestCodes_QuestAcceptedSuccessfully
	res.Message = "Quest accepted successfully"
	return res
}
