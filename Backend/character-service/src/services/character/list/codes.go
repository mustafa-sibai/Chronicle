package list

type ListCharactersCodes int

const (
	ListCharactersCodes_Unknown ListCharactersCodes = iota - 1
	ListCharactersCodes_Success
	ListCharactersCodes_FailedToListCharacters
	ListCharactersCodes_InvalidInput
	ListCharactersCodes_Unauthorized
)
