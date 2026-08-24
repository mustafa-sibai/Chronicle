package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var client *mongo.Client

func Connect(mongodbURI string) {
	clientOptions := options.Client().ApplyURI(mongodbURI)
	var err error
	client, err = mongo.Connect(clientOptions)

	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")

	err = client.Ping(context.Background(), nil)

	if err != nil {
		fmt.Println("Error pinging MongoDB:", err)
	} else {
		fmt.Println("Pinged MongoDB successfully!")
	}
}

func GetMongoDBClient() *mongo.Client {
	return client
}
