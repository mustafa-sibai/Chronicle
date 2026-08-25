package collections

import (
	"context"

	"github.com/mustafa-sibai/chronicle/backend-lib/database/mongodb"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetRefreshTokensCollection() *mongo.Collection {
	return mongodb.GetMongoDBClient().Database("chronicle").Collection("refreshTokens")
}

func IndexRefreshToken(ctx context.Context) error {
	_, err := GetRefreshTokensCollection().Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "expiresAt", Value: 1}},
			Options: options.Index().SetExpireAfterSeconds(0),
		},
		{
			Keys:    bson.D{{Key: "tokenHash", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	})
	return err
}
