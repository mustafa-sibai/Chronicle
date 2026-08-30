package get

import (
	"errors"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type GetItemTemplateBody struct {
	ItemTemplateID string `json:"itemTemplateId"`
}

func (b *GetItemTemplateBody) Normalize() {}

func (b GetItemTemplateBody) Validate() error {
	if b.ItemTemplateID == "" {
		return errors.New("itemTemplateId is required")
	}
	return nil
}

type GetItemTemplateRequest = common.Request[GetItemTemplateBody]

type GetItemTemplateResponse struct {
	common.BaseResponse
	StatusCode   GetItemTemplateCodes `json:"statusCode"`
	ItemTemplate any                  `json:"itemTemplate,omitempty"`
}
