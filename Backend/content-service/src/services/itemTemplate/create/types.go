package create

import (
	"errors"
	"strings"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/game"
)

type CreateItemTemplateBody struct {
	Name      string        `json:"name"`
	ItemType  game.ItemType `json:"itemType"`
	Stats     *game.Stats   `json:"stats,omitempty"`
	MaxStacks int           `json:"maxStacks"`
	Capacity  int           `json:"capacity"`
}

func (b *CreateItemTemplateBody) Normalize() {
	b.Name = strings.TrimSpace(b.Name)
}

func (b CreateItemTemplateBody) Validate() error {
	if b.Name == "" {
		return errors.New("name is required")
	}
	if !b.ItemType.Valid() {
		return errors.New("itemType is invalid")
	}
	if b.MaxStacks < 1 {
		return errors.New("maxStacks must be at least 1")
	}

	isEquipment := b.ItemType == game.ItemType_Weapon || b.ItemType == game.ItemType_Armor
	if isEquipment && b.Stats == nil {
		return errors.New("stats is required for weapons and armor")
	}
	if !isEquipment && b.Stats != nil {
		return errors.New("stats is only valid for weapons and armor")
	}

	if b.ItemType == game.ItemType_Container && b.Capacity <= 0 {
		return errors.New("capacity must be greater than 0 for a container")
	}
	if b.ItemType != game.ItemType_Container && b.Capacity != 0 {
		return errors.New("capacity is only valid for a container")
	}
	return nil
}

type CreateItemTemplateRequest = common.Request[CreateItemTemplateBody]

type CreateItemTemplateResponse struct {
	common.BaseResponse
	StatusCode CreateItemTemplateCodes `json:"statusCode"`
	TemplateID string                  `json:"templateId,omitempty"`
}
