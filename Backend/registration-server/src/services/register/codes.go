package register

type RegisterCode int

const (
	RegisterCodeUnknown RegisterCode = iota - 1
	RegisterCodeUserRegisteredSuccessfully
	RegisterCodeFailedToRegisterUser
	RegisterCodeEmailTaken
	RegisterCodeUsernameTaken
	RegisterCodeInvalidInput
)
