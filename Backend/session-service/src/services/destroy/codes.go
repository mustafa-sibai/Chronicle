package destroy

type DestroySessionCodes int

const (
	DestroySessionCodes_Unknown DestroySessionCodes = iota - 1
	DestroySessionCodes_SessionDestroyedSuccessfully
	DestroySessionCodes_FailedToDestroySession
	DestroySessionCodes_InvalidInput
)
