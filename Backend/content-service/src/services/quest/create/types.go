package create

import (
	"errors"
	"strings"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type QuestObjective struct {
	Description     string `json:"description"`
	CompletionCount int    `json:"completionCount"`
}

type QuestReward struct {
	ItemTemplateIDs       []string `json:"itemTemplateIds"`
	ChoiceItemTemplateIDs []string `json:"choiceItemTemplateIds"`
	Experience            int      `json:"experience"`
}

type CreateQuestBody struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Objective   QuestObjective `json:"objective"`
	Reward      QuestReward    `json:"reward"`
}

func (b *CreateQuestBody) Normalize() {
	b.Title = strings.TrimSpace(b.Title)
	b.Description = strings.TrimSpace(b.Description)
	b.Objective.Description = strings.TrimSpace(b.Objective.Description)
}

func (b CreateQuestBody) Validate() error {
	if b.Title == "" {
		return errors.New("title is required")
	}
	if b.Description == "" {
		return errors.New("description is required")
	}
	if b.Objective.Description == "" {
		return errors.New("objective.description is required")
	}
	if b.Objective.CompletionCount < 1 {
		return errors.New("objective.completionCount must be at least 1")
	}
	if b.Reward.Experience < 0 {
		return errors.New("reward.experience must be at least 0")
	}
	return nil
}

type CreateQuestRequest = common.Request[CreateQuestBody]

type CreateQuestResponse struct {
	common.BaseResponse
	StatusCode CreateQuestCodes `json:"statusCode"`
	QuestID    string           `json:"questId,omitempty"`
}
