package types

import (
	"time"

	"github.com/mustafa-sibai/chronicle/backend-lib/game"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ItemTemplate struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string        `bson:"name" json:"name"`
	ItemType  game.ItemType `bson:"itemType" json:"itemType"`
	MaxStacks int           `bson:"maxStacks" json:"maxStacks"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
}

type ContainerItemTemplate struct {
	ItemTemplate `bson:",inline"`
	Capacity     int `bson:"capacity,omitempty" json:"capacity,omitempty"`
}

type EquipmentItemTemplate struct {
	ItemTemplate `bson:",inline"`
	Stats        game.Stats `bson:"stats" json:"stats"`
}

// DecodeItemTemplate decodes raw BSON into the concrete ItemTemplate variant
// matching its itemType - ContainerItemTemplate, EquipmentItemTemplate, or
// plain ItemTemplate.
func DecodeItemTemplate(raw bson.Raw) (any, error) {
	var probe struct {
		ItemType game.ItemType `bson:"itemType"`
	}
	if err := bson.Unmarshal(raw, &probe); err != nil {
		return nil, err
	}

	switch probe.ItemType {
	case game.ItemType_Container:
		var itemTemplate ContainerItemTemplate
		err := bson.Unmarshal(raw, &itemTemplate)
		return itemTemplate, err
	case game.ItemType_Weapon, game.ItemType_Armor:
		var itemTemplate EquipmentItemTemplate
		err := bson.Unmarshal(raw, &itemTemplate)
		return itemTemplate, err
	default:
		var itemTemplate ItemTemplate
		err := bson.Unmarshal(raw, &itemTemplate)
		return itemTemplate, err
	}
}
