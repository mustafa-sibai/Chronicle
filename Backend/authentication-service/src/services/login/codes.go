package login

type LoginCodes int

const (
	LoginCodes_Unknown LoginCodes = iota - 1
	LoginCodes_UserLoggedInSuccessfully
	LoginCodes_FailedToLogInUser
	LoginCodes_InvalidCredentials
	LoginCodes_InvalidInput
)
