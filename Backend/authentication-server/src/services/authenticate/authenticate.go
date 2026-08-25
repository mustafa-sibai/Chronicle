package authenticate

import (
	"context"
	"errors"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	models "github.com/mustafa-sibai/chronicle/backend-lib/types"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

func AuthenticateAccount(ctx context.Context, req AuthenticateRequest) *AuthenticateResponse {
	config := common.GetConfig()

	res := &AuthenticateResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodeAuthenticate,
		},
		StatusCode: AuthenticateCodeUnknown,
	}

	req.Normalize()

	if err := req.Validate(); err != nil {
		res.HttpCode = common.HttpCodeBadRequest
		res.StatusCode = AuthenticateCodeInvalidInput
		res.Message = err.Error()
		return res
	}

	accounts := collections.GetAccountsCollection()

	var account models.Account
	err := accounts.FindOne(ctx, bson.M{"email": req.Email}).Decode(&account)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			res.HttpCode = common.HttpCodeUnauthorized
			res.StatusCode = AuthenticateCodeInvalidCredentials
			res.Message = "Invalid credentials"
			return res
		}
		res.HttpCode = common.HttpCodeInternalServerError
		res.Message = "Failed to authenticate account"
		return res
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password)); err != nil {
		res.HttpCode = common.HttpCodeUnauthorized
		res.StatusCode = AuthenticateCodeInvalidCredentials
		res.Message = "Invalid credentials"
		return res
	}

	res.HttpCode = common.HttpCodeOK
	res.StatusCode = AuthenticateCodeUserAuthenticatedSuccessfully
	res.Message = "Account authenticated successfully"
	return res
}
