package store

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync" // Import the sync package for synchronization

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	clientOnce sync.Once
)

func ConnectDB() *mongo.Client {
	clientOnce.Do(func() {
		connectionString := os.Getenv("MONGO_URI")
		clientOptions := options.Client().ApplyURI(connectionString)

		c, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		if err := c.Ping(context.Background(), nil); err != nil {
			log.Fatal("Could not connect to MongoDB:", err)
		}
		fmt.Println("Connected to MongoDB")
		client = c
	})

	return client
}

func OpenCollection(collectionName string) *mongo.Collection {
	client := ConnectDB()
	dbName := os.Getenv("DB_NAME")

	return client.Database(dbName).Collection(collectionName)
}
