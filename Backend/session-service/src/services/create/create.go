package create

import (
	"context"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/database/valkeydb"
	"github.com/mustafa-sibai/chronicle/backend-lib/session"
	"github.com/mustafa-sibai/chronicle/backend-lib/token"

	"github.com/valkey-io/valkey-go"
)

func CreateSession(ctx context.Context, req CreateSessionRequest) *CreateSessionResponse {
	config := common.GetConfig()

	res := &CreateSessionResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_CreateSession,
		},
		StatusCode: CreateSessionCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = CreateSessionCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	valkeyClient := valkeydb.GetValkeyClient()

	// GETDEL is atomic: a session exchange code can only ever be consumed once.
	getDelCmd := valkeyClient.B().Getdel().Key(session.SessionExchangeCodeKey(req.Body.SessionExchangeCode)).Build()
	accountID, err := valkeyClient.Do(ctx, getDelCmd).ToString()
	if err != nil {
		if valkey.IsValkeyNil(err) {
			res.HttpCode = common.HttpCodes_Unauthorized
			res.StatusCode = CreateSessionCodes_InvalidSessionExchangeCode
			res.Message = "Invalid or expired session exchange code"
			return res
		}
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to create session"
		return res
	}

	sessionID, err := token.Generate()
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to create session"
		return res
	}

	setCmd := valkeyClient.B().Set().Key(session.SessionKey(sessionID)).Value(accountID).Ex(session.SessionTTL).Build()
	if err := valkeyClient.Do(ctx, setCmd).Error(); err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to create session"
		return res
	}

	res.HttpCode = common.HttpCodes_Created
	res.StatusCode = CreateSessionCodes_SessionCreatedSuccessfully
	res.Message = "Session created successfully"
	res.SessionID = sessionID
	return res
}
