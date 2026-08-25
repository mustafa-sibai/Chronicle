package email

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

func UpdateEmail(ctx context.Context, req UpdateEmailRequest) *UpdateEmailResponse {
	config := common.GetConfig()

	res := &UpdateEmailResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_UpdateEmail,
		},
		StatusCode: UpdateEmailCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = UpdateEmailCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	accountID, ok, err := session.ResolveAccountID(ctx, req.Head.SessionID)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to update email"
		return res
	}
	if !ok {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = UpdateEmailCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	objectID, err := bson.ObjectIDFromHex(accountID)
	if err != nil {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = UpdateEmailCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	accountsDBCollection := collections.GetAccountsCollection()

	err = accountsDBCollection.FindOne(ctx, bson.M{"email": req.Body.NewEmail, "_id": bson.M{"$ne": objectID}}).Err()
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to update email"
		return res
	}
	if err == nil {
		res.HttpCode = common.HttpCodes_Conflict
		res.StatusCode = UpdateEmailCodes_EmailTaken
		res.Message = "Email already registered"
		return res
	}

	update := bson.M{"$set": bson.M{"email": req.Body.NewEmail, "updatedAt": time.Now().UTC()}}
	if _, err := accountsDBCollection.UpdateByID(ctx, objectID, update); err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to update email"
		return res
	}

	res.HttpCode = common.HttpCodes_OK
	res.StatusCode = UpdateEmailCodes_EmailUpdatedSuccessfully
	res.Message = "Email updated successfully"
	return res
}
