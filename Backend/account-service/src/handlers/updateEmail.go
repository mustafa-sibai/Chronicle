package handlers

import (
	"net/http"

	"github.com/mustafa-sibai/chronicle/account-service/services/email"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

func UpdateEmailHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[email.UpdateEmailBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &email.UpdateEmailResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_UpdateEmail,
				Message:         "invalid request body",
			},
			StatusCode: email.UpdateEmailCodes_InvalidInput,
		})
		return
	}

	res := email.UpdateEmail(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}
