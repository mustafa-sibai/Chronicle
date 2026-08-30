package delete

type DeleteQuestCodes int

const (
	DeleteQuestCodes_Unknown DeleteQuestCodes = iota - 1
	DeleteQuestCodes_QuestDeletedSuccessfully
	DeleteQuestCodes_FailedToDeleteQuest
	DeleteQuestCodes_QuestNotFound
	DeleteQuestCodes_InvalidInput
)
