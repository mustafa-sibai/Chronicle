package collections

import (
	"context"

	"github.com/mustafa-sibai/chronicle/backend-lib/database/mongodb"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetInventoriesCollection() *mongo.Collection {
	return mongodb.GetMongoDBClient().Database("chronicle").Collection("inventories")
}

func IndexInventories(ctx context.Context) error {
	_, err := GetInventoriesCollection().Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "characterId", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	return err
}
