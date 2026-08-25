package register

import (
	"errors"
	"strings"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type RegisterBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (b *RegisterBody) Normalize() {
	b.Username = strings.TrimSpace(b.Username)
	b.Email = strings.TrimSpace(b.Email)
}

func (b RegisterBody) Validate() error {
	if b.Username == "" {
		return errors.New("username is required")
	}
	if b.Email == "" {
		return errors.New("email is required")
	}
	if b.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

type RegisterRequest = common.Request[RegisterBody]

type RegisterResponse struct {
	common.BaseResponse
	StatusCode RegisterCodes `json:"statusCode"`
}
