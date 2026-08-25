package email

import (
	"errors"
	"strings"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type UpdateEmailBody struct {
	NewEmail string `json:"newEmail"`
}

func (b *UpdateEmailBody) Normalize() {
	b.NewEmail = strings.TrimSpace(b.NewEmail)
}

func (b UpdateEmailBody) Validate() error {
	if b.NewEmail == "" {
		return errors.New("newEmail is required")
	}
	return nil
}

type UpdateEmailRequest = common.Request[UpdateEmailBody]

type UpdateEmailResponse struct {
	common.BaseResponse
	StatusCode UpdateEmailCodes `json:"statusCode"`
}
