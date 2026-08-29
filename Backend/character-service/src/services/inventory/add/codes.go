package add

type AddItemCodes int

const (
	AddItemCodes_Unknown AddItemCodes = iota - 1
	AddItemCodes_ItemAddedSuccessfully
	AddItemCodes_FailedToAddItem
	AddItemCodes_CharacterNotFound
	AddItemCodes_TemplateNotFound
	AddItemCodes_InvalidInput
	AddItemCodes_InventoryFull
	AddItemCodes_Unauthorized
)
