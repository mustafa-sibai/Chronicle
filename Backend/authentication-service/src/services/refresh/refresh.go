package refresh

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
)

func RefreshToken(ctx context.Context, req RefreshRequest) *RefreshResponse {
	config := common.GetConfig()

	res := &RefreshResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_Refresh,
		},
		StatusCode: RefreshCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = RefreshCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	refreshTokenDBCollection := collections.GetRefreshTokensCollection()
	incomingHash := token.Hash(req.Body.RefreshToken)

	var stored types.RefreshToken
	err := refreshTokenDBCollection.FindOne(ctx, bson.M{"tokenHash": incomingHash}).Decode(&stored)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			res.HttpCode = common.HttpCodes_Unauthorized
			res.StatusCode = RefreshCodes_InvalidToken
			res.Message = "Invalid refresh token"
			return res
		}
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to refresh token"
		return res
	}

	if time.Now().UTC().After(stored.ExpiresAt) {
		refreshTokenDBCollection.DeleteOne(ctx, bson.M{"_id": stored.ID})
		res.HttpCode = common.HttpCodes_Unauthorized
		res.StatusCode = RefreshCodes_InvalidToken
		res.Message = "Invalid refresh token"
		return res
	}

	// Rotate: issue a new refresh token and invalidate the old one.
	newRefreshRaw, err := token.Generate()
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to refresh token"
		return res
	}

	now := time.Now().UTC()
	newRefresh := &types.RefreshToken{
		TokenHash: token.Hash(newRefreshRaw),
		AccountID: stored.AccountID,
		ExpiresAt: now.Add(types.RefreshTokenTTL),
		CreatedAt: now,
	}

	if _, err := refreshTokenDBCollection.InsertOne(ctx, newRefresh); err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to refresh token"
		return res
	}

	if _, err := refreshTokenDBCollection.DeleteOne(ctx, bson.M{"_id": stored.ID}); err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to refresh token"
		return res
	}

	sessionExchangeCode, err := token.Generate()
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to refresh token"
		return res
	}

	valkeyClient := valkeydb.GetValkeyClient()
	setCmd := valkeyClient.B().Set().
		Key(session.SessionExchangeCodeKey(sessionExchangeCode)).
		Value(stored.AccountID.Hex()).
		Ex(session.SessionExchangeCodeTTL).
		Build()
	if err := valkeyClient.Do(ctx, setCmd).Error(); err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to refresh token"
		return res
	}

	res.HttpCode = common.HttpCodes_OK
	res.StatusCode = RefreshCodes_TokenRefreshedSuccessfully
	res.Message = "Token refreshed successfully"
	res.SessionExchangeCode = sessionExchangeCode
	res.RefreshToken = newRefreshRaw
	return res
}
