package collections

import (
	"context"

	"github.com/mustafa-sibai/chronicle/backend-lib/database/mongodb"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetCharactersCollection() *mongo.Collection {
	return mongodb.GetMongoDBClient().Database("chronicle").Collection("characters")
}

func IndexCharacters(ctx context.Context) error {
	_, err := GetCharactersCollection().Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{{Key: "accountId", Value: 1}},
		},
	})
	return err
}
