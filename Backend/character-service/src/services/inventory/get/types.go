package get

import (
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/types"
)

type GetInventoryBody struct{}

func (b *GetInventoryBody) Normalize() {}

func (b GetInventoryBody) Validate() error {
	return nil
}

type GetInventoryRequest = common.Request[GetInventoryBody]

type GetInventoryResponse struct {
	common.BaseResponse
	StatusCode GetInventoryCodes `json:"statusCode"`
	Inventory  *types.Inventory  `json:"inventory,omitempty"`
}
