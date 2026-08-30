package choose

import (
	"errors"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type ChooseCharacterBody struct {
	CharacterID string `json:"characterId"`
}

func (b *ChooseCharacterBody) Normalize() {}

func (b ChooseCharacterBody) Validate() error {
	if b.CharacterID == "" {
		return errors.New("characterId is required")
	}
	return nil
}

type ChooseCharacterRequest = common.Request[ChooseCharacterBody]

type ChooseCharacterResponse struct {
	common.BaseResponse
	StatusCode ChooseCharacterCodes `json:"statusCode"`
}
