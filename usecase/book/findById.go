package book

import (
	"context"
	"time"

	"github.com/agungsptr/go-redis/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindById(db *mongo.Client, id primitive.ObjectID) (models.Book, error) {
	var (
		result models.Book
		ctx    context.Context
		coll   *mongo.Collection
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll = db.Database("library").Collection("books")

	err := coll.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}
