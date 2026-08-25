package create

import (
	"errors"
	"strings"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type CreateSessionBody struct {
	SessionExchangeCode string `json:"sessionExchangeCode"`
}

func (b *CreateSessionBody) Normalize() {
	b.SessionExchangeCode = strings.TrimSpace(b.SessionExchangeCode)
}

func (b CreateSessionBody) Validate() error {
	if b.SessionExchangeCode == "" {
		return errors.New("sessionExchangeCode is required")
	}
	return nil
}

type CreateSessionRequest = common.Request[CreateSessionBody]

type CreateSessionResponse struct {
	common.BaseResponse
	StatusCode CreateSessionCodes `json:"statusCode"`
	SessionID  string             `json:"sessionId,omitempty"`
}
