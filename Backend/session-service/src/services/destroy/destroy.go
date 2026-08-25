package destroy

import (
	"context"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/database/valkeydb"
	"github.com/mustafa-sibai/chronicle/backend-lib/session"
)

func DestroySession(ctx context.Context, req DestroySessionRequest) *DestroySessionResponse {
	config := common.GetConfig()

	res := &DestroySessionResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_DestroySession,
		},
		StatusCode: DestroySessionCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = DestroySessionCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	valkeyClient := valkeydb.GetValkeyClient()
	delCmd := valkeyClient.B().Del().Key(session.SessionKey(req.Body.SessionID)).Build()
	if err := valkeyClient.Do(ctx, delCmd).Error(); err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to destroy session"
		return res
	}

	res.HttpCode = common.HttpCodes_OK
	res.StatusCode = DestroySessionCodes_SessionDestroyedSuccessfully
	res.Message = "Session destroyed successfully"
	return res
}
