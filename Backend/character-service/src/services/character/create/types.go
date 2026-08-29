package create

import (
	"errors"
	"strconv"
	"strings"
	"unicode"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/game"
)

type CreateCharacterBody struct {
	Name       string          `json:"name"`
	Race       game.Race       `json:"race"`
	Class      game.Class      `json:"class"`
	Gender     game.Gender     `json:"gender"`
	Appearance game.Appearance `json:"appearance"`
}

func (b *CreateCharacterBody) Normalize() {
	b.Name = strings.TrimSpace(b.Name)
}

func (b CreateCharacterBody) Validate() error {
	if b.Name == "" {
		return errors.New("name is required")
	}
	if len(b.Name) < game.MinCharacterNameLength || len(b.Name) > game.MaxCharacterNameLength {
		return errors.New("name must be between " + strconv.Itoa(game.MinCharacterNameLength) + " and " + strconv.Itoa(game.MaxCharacterNameLength) + " characters")
	}
	for _, r := range b.Name {
		if !unicode.IsLetter(r) {
			return errors.New("name must only contain letters")
		}
	}
	if !b.Race.Valid() {
		return errors.New("race is invalid")
	}
	if !b.Gender.Valid() {
		return errors.New("gender is invalid")
	}
	if !b.Appearance.Valid() {
		return errors.New("appearance is invalid")
	}
	return nil
}

type CreateCharacterRequest = common.Request[CreateCharacterBody]

type CreateCharacterResponse struct {
	common.BaseResponse
	StatusCode  CreateCharacterCodes `json:"statusCode"`
	CharacterID string               `json:"characterId,omitempty"`
}
