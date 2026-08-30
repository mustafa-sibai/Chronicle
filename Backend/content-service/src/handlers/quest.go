package handlers

import (
	"net/http"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/content-service/services/quest/create"
	"github.com/mustafa-sibai/chronicle/content-service/services/quest/delete"
	"github.com/mustafa-sibai/chronicle/content-service/services/quest/get"
	"github.com/mustafa-sibai/chronicle/content-service/services/quest/list"
)

func CreateQuestHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[create.CreateQuestBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &create.CreateQuestResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_CreateQuest,
				Message:         "invalid request body",
			},
			StatusCode: create.CreateQuestCodes_InvalidInput,
		})
		return
	}

	res := create.CreateQuest(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}

func GetQuestHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[get.GetQuestBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &get.GetQuestResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_GetQuest,
				Message:         "invalid request body",
			},
			StatusCode: get.GetQuestCodes_InvalidInput,
		})
		return
	}

	res := get.GetQuest(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}

func ListQuestsHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[list.ListQuestsBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &list.ListQuestsResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_ListQuests,
				Message:         "invalid request body",
			},
			StatusCode: list.ListQuestsCodes_InvalidInput,
		})
		return
	}

	res := list.ListQuests(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}

func DeleteQuestHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[delete.DeleteQuestBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &delete.DeleteQuestResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_DeleteQuest,
				Message:         "invalid request body",
			},
			StatusCode: delete.DeleteQuestCodes_InvalidInput,
		})
		return
	}

	res := delete.DeleteQuest(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}
