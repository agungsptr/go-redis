package book

import (
	"context"
	"time"

	"github.com/agungsptr/go-redis/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create(db *mongo.Client, data models.Book) (models.Book, error) {
	var (
		result models.Book
		ctx    context.Context
		coll   *mongo.Collection
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll = db.Database("library").Collection("books")

	resInsert, errInsert := coll.InsertOne(ctx, data)
	if errInsert != nil {
		return result, errInsert
	}

	errFind := coll.FindOne(ctx, bson.M{"_id": resInsert.InsertedID}).Decode(&result)
	if errFind != nil {
		return result, errFind
	}

	return result, nil
}
