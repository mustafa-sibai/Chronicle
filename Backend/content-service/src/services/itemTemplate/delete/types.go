package delete

import (
	"errors"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type DeleteItemTemplateBody struct {
	TemplateID string `json:"templateId"`
}

func (b *DeleteItemTemplateBody) Normalize() {}

func (b DeleteItemTemplateBody) Validate() error {
	if b.TemplateID == "" {
		return errors.New("templateId is required")
	}
	return nil
}

type DeleteItemTemplateRequest = common.Request[DeleteItemTemplateBody]

type DeleteItemTemplateResponse struct {
	common.BaseResponse
	StatusCode DeleteItemTemplateCodes `json:"statusCode"`
}
