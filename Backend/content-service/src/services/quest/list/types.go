package list

import (
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/backend-lib/types"
)

type ListQuestsBody struct{}

func (b *ListQuestsBody) Normalize() {}

func (b ListQuestsBody) Validate() error {
	return nil
}

type ListQuestsRequest = common.Request[ListQuestsBody]

type ListQuestsResponse struct {
	common.BaseResponse
	StatusCode ListQuestsCodes `json:"statusCode"`
	Quests     []types.Quest   `json:"quests,omitempty"`
}
