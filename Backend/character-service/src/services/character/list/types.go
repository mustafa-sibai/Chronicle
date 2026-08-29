package list

import (
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/types"
)

type ListCharactersBody struct{}

func (b *ListCharactersBody) Normalize() {}

func (b ListCharactersBody) Validate() error {
	return nil
}

type ListCharactersRequest = common.Request[ListCharactersBody]

type ListCharactersResponse struct {
	common.BaseResponse
	StatusCode ListCharactersCodes `json:"statusCode"`
	Characters []types.Character   `json:"characters,omitempty"`
}
