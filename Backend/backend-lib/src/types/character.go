package types

import (
	"time"

	"github.com/mustafa-sibai/chronicle/backend-lib/game"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Character struct {
	ID         bson.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	AccountID  bson.ObjectID   `bson:"accountId" json:"accountId"`
	Name       string          `bson:"name" json:"name"`
	Race       game.Race       `bson:"race" json:"race"`
	Class      game.Class      `bson:"class" json:"class"`
	Gender     game.Gender     `bson:"gender" json:"gender"`
	Appearance game.Appearance `bson:"appearance" json:"appearance"`
	CreatedAt  time.Time       `bson:"createdAt" json:"createdAt"`
}
