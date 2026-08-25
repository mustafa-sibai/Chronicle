package authenticate

type AuthenticateCode int

const (
	AuthenticateCodeUnknown AuthenticateCode = iota - 1
	AuthenticateCodeUserAuthenticatedSuccessfully
	AuthenticateCodeFailedToAuthenticateUser
	AuthenticateCodeInvalidCredentials
	AuthenticateCodeInvalidInput
)
