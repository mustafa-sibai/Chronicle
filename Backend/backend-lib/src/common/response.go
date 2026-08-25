package common

type BaseResponse struct {
	ApplicationName ApplicationName `json:"applicationName"`
	EnvironmentType EnvironmentType `json:"environmentType"`
	HttpCode        HttpCode        `json:"httpCode"`
	ResponseCode    ResponseCode    `json:"responseCode"`
	Message         string          `json:"message"`
}
