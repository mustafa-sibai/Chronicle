package handlers

import (
	"net/http"

	"github.com/mustafa-sibai/chronicle/account-service/services/password"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

func UpdatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[password.UpdatePasswordBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &password.UpdatePasswordResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_UpdatePassword,
				Message:         "invalid request body",
			},
			StatusCode: password.UpdatePasswordCodes_InvalidInput,
		})
		return
	}

	res := password.UpdatePassword(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}
