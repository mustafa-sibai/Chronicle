package destroy

import (
	"errors"
	"strings"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type DestroySessionBody struct {
	SessionID string `json:"sessionId"`
}

func (b *DestroySessionBody) Normalize() {
	b.SessionID = strings.TrimSpace(b.SessionID)
}

func (b DestroySessionBody) Validate() error {
	if b.SessionID == "" {
		return errors.New("sessionId is required")
	}
	return nil
}

type DestroySessionRequest = common.Request[DestroySessionBody]

type DestroySessionResponse struct {
	common.BaseResponse
	StatusCode DestroySessionCodes `json:"statusCode"`
}
