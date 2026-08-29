package game

type ItemType string

const (
	ItemType_Weapon        ItemType = "weapon"
	ItemType_Armor         ItemType = "armor"
	ItemType_Container     ItemType = "container"
	ItemType_Consumable    ItemType = "consumable"
	ItemType_QuestItem     ItemType = "quest_item"
	ItemType_Miscellaneous ItemType = "miscellaneous"
)

func (t ItemType) Valid() bool {
	switch t {
	case ItemType_Weapon, ItemType_Armor, ItemType_Container, ItemType_Consumable, ItemType_QuestItem, ItemType_Miscellaneous:
		return true
	}
	return false
}
