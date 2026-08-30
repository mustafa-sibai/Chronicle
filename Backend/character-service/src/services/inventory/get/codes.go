package get

type GetInventoryCodes int

const (
	GetInventoryCodes_Unknown GetInventoryCodes = iota - 1
	GetInventoryCodes_Success
	GetInventoryCodes_FailedToGetInventory
	GetInventoryCodes_CharacterNotFound
	GetInventoryCodes_InvalidInput
	GetInventoryCodes_Unauthorized
	GetInventoryCodes_NoCharacterActive
)
