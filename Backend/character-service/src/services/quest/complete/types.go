package complete

import (
	"errors"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type CompleteQuestBody struct {
	QuestID                    string `json:"questId"`
	ChosenRewardItemTemplateID string `json:"chosenRewardItemTemplateId,omitempty"`
}

func (b *CompleteQuestBody) Normalize() {}

func (b CompleteQuestBody) Validate() error {
	if b.QuestID == "" {
		return errors.New("questId is required")
	}
	return nil
}

type CompleteQuestRequest = common.Request[CompleteQuestBody]

type CompleteQuestResponse struct {
	common.BaseResponse
	StatusCode CompleteQuestCodes `json:"statusCode"`
}
