package username

import (
	"context"
	"errors"
	"time"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/session"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func UpdateUsername(ctx context.Context, req UpdateUsernameRequest) *UpdateUsernameResponse {
	config := common.GetConfig()

	res := &UpdateUsernameResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_UpdateUsername,
		},
		StatusCode: UpdateUsernameCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = UpdateUsernameCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	accountID, ok, err := session.ResolveAccountID(ctx, req.Head.SessionID)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to update username"
		return res
	}
	if !ok {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = UpdateUsernameCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	objectID, err := bson.ObjectIDFromHex(accountID)
	if err != nil {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = UpdateUsernameCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	accountsDBCollection := collections.GetAccountsCollection()

	err = accountsDBCollection.FindOne(ctx, bson.M{"username": req.Body.NewUsername, "_id": bson.M{"$ne": objectID}}).Err()
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to update username"
		return res
	}
	if err == nil {
		res.HttpCode = common.HttpCodes_Conflict
		res.StatusCode = UpdateUsernameCodes_UsernameTaken
		res.Message = "Username already taken"
		return res
	}

	update := bson.M{"$set": bson.M{"username": req.Body.NewUsername, "updatedAt": time.Now().UTC()}}
	if _, err := accountsDBCollection.UpdateByID(ctx, objectID, update); err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to update username"
		return res
	}

	res.HttpCode = common.HttpCodes_OK
	res.StatusCode = UpdateUsernameCodes_UsernameUpdatedSuccessfully
	res.Message = "Username updated successfully"
	return res
}
