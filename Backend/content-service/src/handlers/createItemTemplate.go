package handlers

import (
	"net/http"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/content-service/services/itemTemplate/create"
)

func CreateItemTemplateHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[create.CreateItemTemplateBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &create.CreateItemTemplateResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_CreateItemTemplate,
				Message:         "invalid request body",
			},
			StatusCode: create.CreateItemTemplateCodes_InvalidInput,
		})
		return
	}

	res := create.CreateItemTemplate(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}
