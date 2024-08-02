package book

import (
	"context"
	"time"

	"github.com/agungsptr/go-redis/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindAll(db *mongo.Client, limit int64, filters ...bson.M) ([]models.Book, error) {
	var (
		result   []models.Book
		ctx      context.Context
		coll     *mongo.Collection
		filterDb bson.M
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll = db.Database("library").Collection("books")

	// Filter
	if len(filters) > 0 {
		filterDb = bson.M{"$and": filters}
	} else {
		filterDb = bson.M{}
	}

	// Option
	opts := []*options.FindOptions{
		options.Find().SetLimit(limit),
	}

	cursor, err := coll.Find(ctx, filterDb, opts...)
	if err != nil {
		return result, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var book models.Book
		err := cursor.Decode(&book)
		if err != nil {
			return nil, err
		}
		result = append(result, book)
	}

	return result, nil
}
