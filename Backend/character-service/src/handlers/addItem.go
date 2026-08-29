package handlers

import (
	"net/http"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/character-service/services/inventory/add"
)

func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[add.AddItemBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &add.AddItemResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_AddItem,
				Message:         "invalid request body",
			},
			StatusCode: add.AddItemCodes_InvalidInput,
		})
		return
	}

	res := add.AddItem(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}
