package delete

import (
	"errors"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
)

type DeleteQuestBody struct {
	QuestID string `json:"questId"`
}

func (b *DeleteQuestBody) Normalize() {}

func (b DeleteQuestBody) Validate() error {
	if b.QuestID == "" {
		return errors.New("questId is required")
	}
	return nil
}

type DeleteQuestRequest = common.Request[DeleteQuestBody]

type DeleteQuestResponse struct {
	common.BaseResponse
	StatusCode DeleteQuestCodes `json:"statusCode"`
}
