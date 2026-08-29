package add

import (
	"errors"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type AddItemBody struct {
	CharacterID string `json:"characterId"`
	TemplateID  string `json:"templateId"`
	Quantity    int    `json:"quantity"`
}

func (b *AddItemBody) Normalize() {}

func (b AddItemBody) Validate() error {
	if b.CharacterID == "" {
		return errors.New("characterId is required")
	}
	if b.TemplateID == "" {
		return errors.New("templateId is required")
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
