package handlers

import (
	"net/http"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/content-service/services/itemTemplate/delete"
)

func DeleteItemTemplateHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[delete.DeleteItemTemplateBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &delete.DeleteItemTemplateResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_DeleteItemTemplate,
				Message:         "invalid request body",
			},
			StatusCode: delete.DeleteItemTemplateCodes_InvalidInput,
		})
		return
	}

	res := delete.DeleteItemTemplate(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}
