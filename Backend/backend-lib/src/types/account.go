package types

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Account struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string        `bson:"username" json:"username"`
	Email     string        `bson:"email" json:"email"`
	Password  string        `bson:"password" json:"-"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt" json:"updatedAt"`
}
