package common

type HttpCode int

const (
	HttpCodeUnknown             HttpCode = -1
	HttpCodeOK                  HttpCode = 200
	HttpCodeCreated             HttpCode = 201
	HttpCodeBadRequest          HttpCode = 400
	HttpCodeUnauthorized        HttpCode = 401
	HttpCodeConflict            HttpCode = 409
	HttpCodeInternalServerError HttpCode = 500
)

type ResponseCode int

const (
	ResponseCodeUnknown      ResponseCode = iota - 1 // -1
	ResponseCodeRegister                             // 0
	ResponseCodeAuthenticate                         // 1
)
