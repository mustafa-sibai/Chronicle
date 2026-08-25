package collections

import (
	"github.com/mustafa-sibai/chronicle/backend-lib/database/mongodb"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetAccountsCollection() *mongo.Collection {
	return mongodb.GetMongoDBClient().Database("chronicle").Collection("accounts")
}
