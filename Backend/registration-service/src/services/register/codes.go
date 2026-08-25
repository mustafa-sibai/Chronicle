package register

type RegisterCodes int

const (
	RegisterCodes_Unknown RegisterCodes = iota - 1
	RegisterCodes_UserRegisteredSuccessfully
	RegisterCodes_FailedToRegisterUser
	RegisterCodes_EmailTaken
	RegisterCodes_UsernameTaken
	RegisterCodes_InvalidInput
)
