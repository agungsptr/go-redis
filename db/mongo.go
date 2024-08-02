package db

import (
	"context"
	"log"
	"time"

	"github.com/agungsptr/go-redis/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoClient() *mongo.Client {
	// Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to mongodb
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Get().MongoUri))
	if err != nil {
		panic(err)
	}
	defer Disconnect(client)

	return client
}

func Disconnect(client *mongo.Client) {
	errMsg := recover()
	if errMsg != nil {
		log.Fatal(errMsg)
	}
}
