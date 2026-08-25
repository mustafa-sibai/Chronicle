package create

type CreateSessionCodes int

const (
	CreateSessionCodes_Unknown CreateSessionCodes = iota - 1
	CreateSessionCodes_SessionCreatedSuccessfully
	CreateSessionCodes_FailedToCreateSession
	CreateSessionCodes_InvalidSessionExchangeCode
	CreateSessionCodes_InvalidInput
)
