package email

type UpdateEmailCodes int

const (
	UpdateEmailCodes_Unknown UpdateEmailCodes = iota - 1
	UpdateEmailCodes_EmailUpdatedSuccessfully
	UpdateEmailCodes_FailedToUpdateEmail
	UpdateEmailCodes_EmailTaken
	UpdateEmailCodes_InvalidInput
	UpdateEmailCodes_Unauthorized
)
