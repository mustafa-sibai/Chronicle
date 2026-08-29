package types

import (
	"github.com/mustafa-sibai/chronicle/backend-lib/game"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Inventory struct {
	ID          bson.ObjectID                `bson:"_id,omitempty" json:"id,omitempty"`
	CharacterID bson.ObjectID                `bson:"characterId" json:"characterId"`
	Bags        [game.InventoryBagSlots]*Bag `bson:"bags" json:"bags"`
}
