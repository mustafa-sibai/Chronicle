package handlers

import (
	"net/http"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/registration-service/services/register"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[register.RegisterBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &register.RegisterResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_Register,
				Message:         "invalid request body",
			},
			StatusCode: register.RegisterCodes_InvalidInput,
		})
		return
	}

	res := register.RegisterAccount(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}
