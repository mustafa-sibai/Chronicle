package handlers

import (
	"net/http"

	"github.com/mustafa-sibai/chronicle/account-service/services/username"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

func UpdateUsernameHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[username.UpdateUsernameBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &username.UpdateUsernameResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_UpdateUsername,
				Message:         "invalid request body",
			},
			StatusCode: username.UpdateUsernameCodes_InvalidInput,
		})
		return
	}

	res := username.UpdateUsername(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}
