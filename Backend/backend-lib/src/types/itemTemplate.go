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
