package accept

import (
	"errors"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type AcceptQuestBody struct {
	QuestID string `json:"questId"`
}

func (b *AcceptQuestBody) Normalize() {}

func (b AcceptQuestBody) Validate() error {
	if b.QuestID == "" {
		return errors.New("questId is required")
	}
	return nil
}

type AcceptQuestRequest = common.Request[AcceptQuestBody]

type AcceptQuestResponse struct {
	common.BaseResponse
	StatusCode AcceptQuestCodes `json:"statusCode"`
}
