package list

import "github.com/mustafa-sibai/chronicle/backend-lib/common"

type ListItemTemplatesBody struct{}

func (b *ListItemTemplatesBody) Normalize() {}

func (b ListItemTemplatesBody) Validate() error {
	return nil
}

type ListItemTemplatesRequest = common.Request[ListItemTemplatesBody]

type ListItemTemplatesResponse struct {
	common.BaseResponse
	StatusCode    ListItemTemplatesCodes `json:"statusCode"`
	ItemTemplates []any                  `json:"itemTemplates,omitempty"`
}
