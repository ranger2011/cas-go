package data

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectString string = "mongodb://localhost:27017"

func CheckDatabase() bool {
	clientOptions := options.Client().ApplyURI(connectString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return false
	}
	return client.Ping(context.TODO(), nil) == nil
}

func GetCollection(dbName string, collectionName string) *mongo.Collection {
	clientOptions := options.Client().ApplyURI(connectString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil
	}
	return client.Database(dbName).Collection(collectionName)
}
