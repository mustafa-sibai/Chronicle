package create

type CreateQuestCodes int

const (
	CreateQuestCodes_Unknown CreateQuestCodes = iota - 1
	CreateQuestCodes_QuestCreatedSuccessfully
	CreateQuestCodes_FailedToCreateQuest
	CreateQuestCodes_InvalidInput
	CreateQuestCodes_RewardItemTemplateNotFound
)
