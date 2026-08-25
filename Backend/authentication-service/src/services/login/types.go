package login

import (
	"errors"
	"strings"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type LoginBody struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberMe bool   `json:"rememberMe"`
}

func (b *LoginBody) Normalize() {
	b.Email = strings.TrimSpace(b.Email)
}

func (b LoginBody) Validate() error {
	if b.Email == "" {
		return errors.New("email is required")
	}
	if b.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

type LoginRequest = common.Request[LoginBody]

type LoginResponse struct {
	common.BaseResponse
	StatusCode          LoginCodes `json:"statusCode"`
	SessionExchangeCode string     `json:"sessionExchangeCode,omitempty"`
	RefreshToken        string     `json:"refreshToken,omitempty"`
}
