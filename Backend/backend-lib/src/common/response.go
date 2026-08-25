package common

import (
	"encoding/json"
	"net/http"
)

type BaseResponse struct {
	ApplicationName ApplicationName `json:"applicationName"`
	EnvironmentType EnvironmentType `json:"environmentType"`
	HttpCode        HttpCodes       `json:"httpCode"`
	ResponseCode    ResponseCodes   `json:"responseCode"`
	Message         string          `json:"message"`
}

func WriteJSON(w http.ResponseWriter, httpCode HttpCodes, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(httpCode))
	json.NewEncoder(w).Encode(v)
}
