package handlers

import (
	"net/http"

	"github.com/mustafa-sibai/chronicle/authentication-service/services/refresh"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[refresh.RefreshBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &refresh.RefreshResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_Refresh,
				Message:         "invalid request body",
			},
			StatusCode: refresh.RefreshCodes_InvalidInput,
		})
		return
	}

	res := refresh.RefreshToken(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}
