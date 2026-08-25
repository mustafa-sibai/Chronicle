package handlers

import (
	"net/http"

	"github.com/mustafa-sibai/chronicle/authentication-service/services/logout"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[logout.LogoutBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &logout.LogoutResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_Logout,
				Message:         "invalid request body",
			},
			StatusCode: logout.LogoutCodes_InvalidInput,
		})
		return
	}

	res := logout.Logout(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}
