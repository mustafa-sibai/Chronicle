package refresh

type RefreshCodes int

const (
	RefreshCodes_Unknown RefreshCodes = iota - 1
	RefreshCodes_TokenRefreshedSuccessfully
	RefreshCodes_FailedToRefreshToken
	RefreshCodes_InvalidToken
	RefreshCodes_InvalidInput
)
