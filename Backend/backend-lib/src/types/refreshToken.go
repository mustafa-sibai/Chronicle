package types

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const RefreshTokenTTL = 14 * 24 * time.Hour

type RefreshToken struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TokenHash string        `bson:"tokenHash" json:"-"`
	AccountID bson.ObjectID `bson:"accountId" json:"accountId"`
	ExpiresAt time.Time     `bson:"expiresAt" json:"expiresAt"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
}
