package refresh

import (
	"errors"
	"strings"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type RefreshBody struct {
	RefreshToken string `json:"refreshToken"`
}

func (b *RefreshBody) Normalize() {
	b.RefreshToken = strings.TrimSpace(b.RefreshToken)
}

func (b RefreshBody) Validate() error {
	if b.RefreshToken == "" {
		return errors.New("refreshToken is required")
	}
	return nil
}

type RefreshRequest = common.Request[RefreshBody]

type RefreshResponse struct {
	common.BaseResponse
	StatusCode          RefreshCodes `json:"statusCode"`
	SessionExchangeCode string       `json:"sessionExchangeCode,omitempty"`
	RefreshToken        string       `json:"refreshToken,omitempty"`
}
