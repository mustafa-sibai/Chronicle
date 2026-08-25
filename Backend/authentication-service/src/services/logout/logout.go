package logout

import (
	"context"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/token"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func Logout(ctx context.Context, req LogoutRequest) *LogoutResponse {
	config := common.GetConfig()

	res := &LogoutResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_Logout,
		},
		StatusCode: LogoutCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = LogoutCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	refreshTokenDBCollection := collections.GetRefreshTokensCollection()
	if _, err := refreshTokenDBCollection.DeleteOne(ctx, bson.M{"tokenHash": token.Hash(req.Body.RefreshToken)}); err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to logout"
		return res
	}

	res.HttpCode = common.HttpCodes_OK
	res.StatusCode = LogoutCodes_LoggedOutSuccessfully
	res.Message = "Logged out successfully"
	return res
}
