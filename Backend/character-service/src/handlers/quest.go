package handlers

import (
	"net/http"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/character-service/services/quest/accept"
	"github.com/mustafa-sibai/chronicle/character-service/services/quest/complete"
	"github.com/mustafa-sibai/chronicle/character-service/services/quest/progress"
)

func AcceptQuestHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[accept.AcceptQuestBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &accept.AcceptQuestResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_AcceptQuest,
				Message:         "invalid request body",
			},
			StatusCode: accept.AcceptQuestCodes_InvalidInput,
		})
		return
	}

	res := accept.AcceptQuest(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}

func UpdateQuestProgressHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[progress.UpdateQuestProgressBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &progress.UpdateQuestProgressResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_UpdateQuestProgress,
				Message:         "invalid request body",
			},
			StatusCode: progress.UpdateQuestProgressCodes_InvalidInput,
		})
		return
	}

	res := progress.UpdateQuestProgress(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}

func CompleteQuestHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[complete.CompleteQuestBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &complete.CompleteQuestResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_CompleteQuest,
				Message:         "invalid request body",
			},
			StatusCode: complete.CompleteQuestCodes_InvalidInput,
		})
		return
	}

	res := complete.CompleteQuest(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}
