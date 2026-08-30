package handlers

import (
	"net/http"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/content-service/services/itemTemplate/create"
	"github.com/mustafa-sibai/chronicle/content-service/services/itemTemplate/delete"
	"github.com/mustafa-sibai/chronicle/content-service/services/itemTemplate/get"
	"github.com/mustafa-sibai/chronicle/content-service/services/itemTemplate/list"
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

func ListItemTemplatesHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[list.ListItemTemplatesBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &list.ListItemTemplatesResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_ListItemTemplates,
				Message:         "invalid request body",
			},
			StatusCode: list.ListItemTemplatesCodes_InvalidInput,
		})
		return
	}

	res := list.ListItemTemplates(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}

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
