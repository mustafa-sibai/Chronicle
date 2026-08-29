package get

type GetItemTemplateCodes int

const (
	GetItemTemplateCodes_Unknown GetItemTemplateCodes = iota - 1
	GetItemTemplateCodes_Success
	GetItemTemplateCodes_FailedToGetItemTemplate
	GetItemTemplateCodes_ItemTemplateNotFound
	GetItemTemplateCodes_InvalidInput
)
