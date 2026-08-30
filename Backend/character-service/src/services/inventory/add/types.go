package add

import (
	"errors"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type AddItemBody struct {
	ItemTemplateID string `json:"itemTemplateId"`
	Quantity       int    `json:"quantity"`
}

func (b *AddItemBody) Normalize() {}

func (b AddItemBody) Validate() error {
	if b.ItemTemplateID == "" {
		return errors.New("itemTemplateId is required")
	}
	if b.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}
	return nil
}

type AddItemRequest = common.Request[AddItemBody]

type AddItemResponse struct {
	common.BaseResponse
	StatusCode AddItemCodes `json:"statusCode"`
}
