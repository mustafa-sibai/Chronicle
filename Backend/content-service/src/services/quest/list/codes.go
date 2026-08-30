package list

type ListQuestsCodes int

const (
	ListQuestsCodes_Unknown ListQuestsCodes = iota - 1
	ListQuestsCodes_Success
	ListQuestsCodes_FailedToListQuests
	ListQuestsCodes_InvalidInput
)
