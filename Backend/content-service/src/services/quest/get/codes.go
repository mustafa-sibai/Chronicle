package get

type GetQuestCodes int

const (
	GetQuestCodes_Unknown GetQuestCodes = iota - 1
	GetQuestCodes_Success
	GetQuestCodes_FailedToGetQuest
	GetQuestCodes_QuestNotFound
	GetQuestCodes_InvalidInput
)
