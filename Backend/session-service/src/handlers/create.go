package handlers

import (
	"net/http"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/session-service/services/create"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[create.CreateSessionBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &create.CreateSessionResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_CreateSession,
				Message:         "invalid request body",
			},
			StatusCode: create.CreateSessionCodes_InvalidInput,
		})
		return
	}

	res := create.CreateSession(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}
