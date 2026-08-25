package logout

type LogoutCodes int

const (
	LogoutCodes_Unknown LogoutCodes = iota - 1
	LogoutCodes_LoggedOutSuccessfully
	LogoutCodes_FailedToLogout
	LogoutCodes_InvalidInput
)
