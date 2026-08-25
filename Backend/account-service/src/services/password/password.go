package password

import (
	"context"
	"errors"
	"time"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/session"
	models "github.com/mustafa-sibai/chronicle/backend-lib/types"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

func UpdatePassword(ctx context.Context, req UpdatePasswordRequest) *UpdatePasswordResponse {
	config := common.GetConfig()

	res := &UpdatePasswordResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_UpdatePassword,
		},
		StatusCode: UpdatePasswordCodes_Unknown,
	}

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = UpdatePasswordCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	accountID, ok, err := session.ResolveAccountID(ctx, req.Head.SessionID)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to update password"
		return res
	}
	if !ok {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = UpdatePasswordCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	objectID, err := bson.ObjectIDFromHex(accountID)
	if err != nil {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = UpdatePasswordCodes_Unauthorized
		res.Message = "Invalid or expired session"
		return res
	}

	accountsDBCollection := collections.GetAccountsCollection()

	var account models.Account
	if err := accountsDBCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&account); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			res.HttpCode = common.HttpCodes_Unauthorized
			res.StatusCode = UpdatePasswordCodes_Unauthorized
			res.Message = "Invalid or expired session"
			return res
		}
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to update password"
		return res
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Body.CurrentPassword)); err != nil {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = UpdatePasswordCodes_InvalidCredentials
		res.Message = "Current password is incorrect"
		return res
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Body.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to update password"
		return res
	}

	update := bson.M{"$set": bson.M{"password": string(hashed), "updatedAt": time.Now().UTC()}}
	if _, err := accountsDBCollection.UpdateByID(ctx, objectID, update); err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to update password"
		return res
	}

	res.HttpCode = common.HttpCodes_OK
	res.StatusCode = UpdatePasswordCodes_PasswordUpdatedSuccessfully
	res.Message = "Password updated successfully"
	return res
}
