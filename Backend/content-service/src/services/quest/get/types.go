package get

import (
	"errors"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/types"
)

type GetQuestBody struct {
	QuestID string `json:"questId"`
}

func (b *GetQuestBody) Normalize() {}

func (b GetQuestBody) Validate() error {
	if b.QuestID == "" {
		return errors.New("questId is required")
	}
	return nil
}

type GetQuestRequest = common.Request[GetQuestBody]

type GetQuestResponse struct {
	common.BaseResponse
	StatusCode GetQuestCodes `json:"statusCode"`
	Quest      *types.Quest  `json:"quest,omitempty"`
}
