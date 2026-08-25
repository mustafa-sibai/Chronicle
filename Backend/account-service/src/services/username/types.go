package username

import (
	"errors"
	"strings"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type UpdateUsernameBody struct {
	NewUsername string `json:"newUsername"`
}

func (b *UpdateUsernameBody) Normalize() {
	b.NewUsername = strings.TrimSpace(b.NewUsername)
}

func (b UpdateUsernameBody) Validate() error {
	if b.NewUsername == "" {
		return errors.New("newUsername is required")
	}
	return nil
}

type UpdateUsernameRequest = common.Request[UpdateUsernameBody]

type UpdateUsernameResponse struct {
	common.BaseResponse
	StatusCode UpdateUsernameCodes `json:"statusCode"`
}
