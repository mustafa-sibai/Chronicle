package progress

import (
	"errors"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type UpdateQuestProgressBody struct {
	QuestID string `json:"questId"`
	Amount  int    `json:"amount"`
}

func (b *UpdateQuestProgressBody) Normalize() {}

func (b UpdateQuestProgressBody) Validate() error {
	if b.QuestID == "" {
		return errors.New("questId is required")
	}
	if b.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}
	return nil
}

type UpdateQuestProgressRequest = common.Request[UpdateQuestProgressBody]

type UpdateQuestProgressResponse struct {
	common.BaseResponse
	StatusCode UpdateQuestProgressCodes `json:"statusCode"`
	Progress   int                      `json:"progress,omitempty"`
}
