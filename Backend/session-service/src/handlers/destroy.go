package handlers

import (
	"net/http"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/session-service/services/destroy"
)

func DestroyHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[destroy.DestroySessionBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &destroy.DestroySessionResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_DestroySession,
				Message:         "invalid request body",
			},
			StatusCode: destroy.DestroySessionCodes_InvalidInput,
		})
		return
	}

	res := destroy.DestroySession(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}
