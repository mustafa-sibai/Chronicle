package login

import (
	"context"
	"errors"
	"time"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/database/valkeydb"
	"github.com/mustafa-sibai/chronicle/backend-lib/session"
	"github.com/mustafa-sibai/chronicle/backend-lib/token"
	"github.com/mustafa-sibai/chronicle/backend-lib/types"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

func LoginAccount(ctx context.Context, req LoginRequest) *LoginResponse {
	config := common.GetConfig()

	res := &LoginResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_Login,
		},
		StatusCode: LoginCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = LoginCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	accountsDBCollection := collections.GetAccountsCollection()

	var account types.Account
	err := accountsDBCollection.FindOne(ctx, bson.M{"email": req.Body.Email}).Decode(&account)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			res.HttpCode = common.HttpCodes_Unauthorized
			res.StatusCode = LoginCodes_InvalidCredentials
			res.Message = "Invalid credentials"
			return res
		}
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to log in"
		return res
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Body.Password)); err != nil {
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = LoginCodes_InvalidCredentials
		res.Message = "Invalid credentials"
		return res
	}

	sessionExchangeCode, err := token.Generate()
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to log in"
		return res
	}

	valkeyClient := valkeydb.GetValkeyClient()
	setCmd := valkeyClient.B().Set().Key(session.SessionExchangeCodeKey(sessionExchangeCode)).Value(account.ID.Hex()).Ex(session.SessionExchangeCodeTTL).Build()
	if err := valkeyClient.Do(ctx, setCmd).Error(); err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to log in"
		return res
	}

	res.SessionExchangeCode = sessionExchangeCode

	if req.Body.RememberMe {
		refreshRaw, err := token.Generate()
		if err != nil {
			res.HttpCode = common.HttpCodes_InternalServerError
			res.Message = "Failed to log in"
			return res
		}

		now := time.Now().UTC()
		refreshToken := &types.RefreshToken{
			TokenHash: token.Hash(refreshRaw),
			AccountID: account.ID,
			ExpiresAt: now.Add(types.RefreshTokenTTL),
			CreatedAt: now,
		}

		if _, err := collections.GetRefreshTokensCollection().InsertOne(ctx, refreshToken); err != nil {
			res.HttpCode = common.HttpCodes_InternalServerError
			res.Message = "Failed to log in"
			return res
		}

		res.RefreshToken = refreshRaw
	}

	res.HttpCode = common.HttpCodes_OK
	res.StatusCode = LoginCodes_UserLoggedInSuccessfully
	res.Message = "Logged in successfully"
	return res
}
