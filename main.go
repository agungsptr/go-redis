package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/agungsptr/go-redis/common"
	"github.com/agungsptr/go-redis/db"
	"github.com/agungsptr/go-redis/models"
	"github.com/agungsptr/go-redis/usecase/book"
	"github.com/redis/go-redis/v9"
)

func main() {
	data := models.Book{
		Title:  "Outlier",
		Author: "Malcom Gladwell",
		Publish: models.BookPublish{
			PublisherName: "Gramedia",
			PublishStatus: "Published",
			PublishDate:   time.Now(),
		},
	}

	mongoClient := db.MongoClient()
	redisClient := db.RedisClient()
	ctx := context.Background()

	// Save book to mongoDB
	saveBook, err := book.Create(mongoClient, data)
	handleError(err)
	fmt.Printf("Save Book to mongoDB: \n%s\n\n", common.JsonPrettyPrint(saveBook))

	// Simulate find book data 10 times
	for i := 1; i <= 10; i++ {
		// Check if book available in redis
		val, err := redisClient.Get(ctx, fmt.Sprintf("book:%s", saveBook.Id)).Result()
		if err != nil {
			// Find book from mongoDB
			findBook, err := book.FindById(mongoClient, saveBook.Id)
			handleError(err)

			// Set book to redis
			setBookToRedis(ctx, redisClient, findBook)

			fmt.Printf("Attempt: %d\nBook retrieved from MongoDB: \n%s\n\n", i, common.JsonPrettyPrint(findBook))
		} else {
			// Deserialize the JSON to a Book struct
			var dataFromRedis models.Book
			err = json.Unmarshal([]byte(val), &dataFromRedis)
			handleError(err)

			fmt.Printf("Attempt: %d\nBook retrieved from Redis: \n%s\n\n", i, common.JsonPrettyPrint(dataFromRedis))
		}
	}
}

func setBookToRedis(ctx context.Context, rc *redis.Client, book models.Book) {
	// Serialize book struct to JSON
	dataJson, err := json.Marshal(book)
	handleError(err)

	// Set book to redis, so can get book data from redis as cache
	err = rc.Set(ctx, fmt.Sprintf("book:%s", book.Id), dataJson, 0).Err()
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
