package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mustafa-sibai/chronicle/authentication-server/services/authenticate"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

func AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	var req authenticate.AuthenticateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		config := common.GetConfig()
		writeJSON(w, &authenticate.AuthenticateResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodeBadRequest,
				ResponseCode:    common.ResponseCodeAuthenticate,
				Message:         "invalid request body",
			},
			StatusCode: authenticate.AuthenticateCodeInvalidInput,
		})
		return
	}

	writeJSON(w, authenticate.AuthenticateAccount(r.Context(), req))
}

func writeJSON(w http.ResponseWriter, res *authenticate.AuthenticateResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.HttpCode))
	json.NewEncoder(w).Encode(res)
}
