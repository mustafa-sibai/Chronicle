package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/registration-server/services/register"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req register.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		config := common.GetConfig()
		writeJSON(w, &register.RegisterResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodeBadRequest,
				ResponseCode:    common.ResponseCodeRegister,
				Message:         "invalid request body",
			},
			StatusCode: register.RegisterCodeInvalidInput,
		})
		return
	}

	writeJSON(w, register.RegisterAccount(r.Context(), req))
}

func writeJSON(w http.ResponseWriter, res *register.RegisterResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.HttpCode))
	json.NewEncoder(w).Encode(res)
}
