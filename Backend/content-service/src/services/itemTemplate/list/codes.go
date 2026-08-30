package list

type ListItemTemplatesCodes int

const (
	ListItemTemplatesCodes_Unknown ListItemTemplatesCodes = iota - 1
	ListItemTemplatesCodes_Success
	ListItemTemplatesCodes_FailedToListItemTemplates
	ListItemTemplatesCodes_InvalidInput
)
