package get

import (
	"errors"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/types"
)

type GetInventoryBody struct {
	CharacterID string `json:"characterId"`
}

func (b *GetInventoryBody) Normalize() {}

func (b GetInventoryBody) Validate() error {
	if b.CharacterID == "" {
		return errors.New("characterId is required")
	}
	return nil
}

type GetInventoryRequest = common.Request[GetInventoryBody]

type GetInventoryResponse struct {
	common.BaseResponse
	StatusCode GetInventoryCodes `json:"statusCode"`
	Inventory  *types.Inventory  `json:"inventory,omitempty"`
}
