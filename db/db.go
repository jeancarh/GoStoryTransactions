package db

import (
	"context"
	"log"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"GoStori/models"
)



func SaveToMongoDB(transactions []models.Transaction) error {
	mongoURL := os.Getenv("mongoUrl")
	dbName := os.Getenv("mongoDb")
	collectionName := os.Getenv("mongoCollection")
	clientOptions := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	collection := client.Database(dbName).Collection(collectionName)

	for _, transaction := range transactions {
		_, err := collection.InsertOne(context.Background(), transaction)
		if err != nil {
			return err
		}
	}

	return nil
}