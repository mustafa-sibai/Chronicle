package handlers

import (
	"net/http"

	"github.com/mustafa-sibai/chronicle/authentication-service/services/login"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[login.LoginBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &login.LoginResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_Login,
				Message:         "invalid request body",
			},
			StatusCode: login.LoginCodes_InvalidInput,
		})
		return
	}

	res := login.LoginAccount(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}
