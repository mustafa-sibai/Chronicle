package authenticate

import (
	"errors"
	"strings"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type AuthenticateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *AuthenticateRequest) Normalize() {
	r.Email = strings.TrimSpace(r.Email)
}

func (r AuthenticateRequest) Validate() error {
	if r.Email == "" {
		return errors.New("email is required")
	}
	if r.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

type AuthenticateResponse struct {
	common.BaseResponse
	StatusCode AuthenticateCode `json:"statusCode"`
}
