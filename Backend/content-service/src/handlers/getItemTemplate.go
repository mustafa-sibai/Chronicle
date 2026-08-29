package handlers

import (
	"net/http"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/content-service/services/itemTemplate/get"
)

func GetItemTemplateHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[get.GetItemTemplateBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &get.GetItemTemplateResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_GetItemTemplate,
				Message:         "invalid request body",
			},
			StatusCode: get.GetItemTemplateCodes_InvalidInput,
		})
		return
	}

	res := get.GetItemTemplate(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}
