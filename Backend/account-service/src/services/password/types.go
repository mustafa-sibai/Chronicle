package password

import (
	"errors"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type UpdatePasswordBody struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

func (b UpdatePasswordBody) Validate() error {
	if b.CurrentPassword == "" {
		return errors.New("currentPassword is required")
	}
	if b.NewPassword == "" {
		return errors.New("newPassword is required")
	}
	return nil
}

type UpdatePasswordRequest = common.Request[UpdatePasswordBody]

type UpdatePasswordResponse struct {
	common.BaseResponse
	StatusCode UpdatePasswordCodes `json:"statusCode"`
}
