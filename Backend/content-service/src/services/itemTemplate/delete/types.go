package delete

import (
	"errors"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type DeleteItemTemplateBody struct {
	ItemTemplateID string `json:"itemTemplateId"`
}

func (b *DeleteItemTemplateBody) Normalize() {}

func (b DeleteItemTemplateBody) Validate() error {
	if b.ItemTemplateID == "" {
		return errors.New("itemTemplateId is required")
	}
	return nil
}

type DeleteItemTemplateRequest = common.Request[DeleteItemTemplateBody]

type DeleteItemTemplateResponse struct {
	common.BaseResponse
	StatusCode DeleteItemTemplateCodes `json:"statusCode"`
}
