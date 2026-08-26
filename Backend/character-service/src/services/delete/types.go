package delete

import (
	"errors"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type DeleteCharacterBody struct {
	CharacterID string `json:"characterId"`
}

func (b *DeleteCharacterBody) Normalize() {}

func (b DeleteCharacterBody) Validate() error {
	if b.CharacterID == "" {
		return errors.New("characterId is required")
	}
	return nil
}

type DeleteCharacterRequest = common.Request[DeleteCharacterBody]

type DeleteCharacterResponse struct {
	common.BaseResponse
	StatusCode DeleteCharacterCodes `json:"statusCode"`
}
