package logout

import (
	"errors"
	"strings"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type LogoutBody struct {
	RefreshToken string `json:"refreshToken"`
}

func (b *LogoutBody) Normalize() {
	b.RefreshToken = strings.TrimSpace(b.RefreshToken)
}

func (b LogoutBody) Validate() error {
	if b.RefreshToken == "" {
		return errors.New("refreshToken is required")
	}
	return nil
}

type LogoutRequest = common.Request[LogoutBody]

type LogoutResponse struct {
	common.BaseResponse
	StatusCode LogoutCodes `json:"statusCode"`
}
