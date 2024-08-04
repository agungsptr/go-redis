package main

import (
	"fmt"
	"time"

	"github.com/agungsptr/go-redis/common"
	"github.com/agungsptr/go-redis/db/mongo"
	"github.com/agungsptr/go-redis/models"
	"github.com/agungsptr/go-redis/usecase/book"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	db := mongo.MongoClient()

	data := models.Book{
		Title:  "Outlier",
		Author: "Malcom Gladwell",
		Publish: models.BookPublish{
			PublisherName: "Gramedia",
			PublishStatus: "Published",
			PublishDate:   time.Now(),
		},
	}

	// Create Book
	createBook, err := book.Create(db, data)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("// Create Book\n%s\n", common.JsonMarshal(createBook))

	// Find Book by Id
	findBook, err := book.FindById(db, createBook.Id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("// Find Book by Id\n%s\n", common.JsonMarshal(findBook))

	// Find All Book
	findAllBook, err := book.FindAll(db, 3,
		bson.M{"title": "Outlier"},
		bson.M{"publish.publish_status": "Published"},
	)
	if err != nil {
		panic((err.Error()))
	}
	fmt.Printf("// Find All Book\n%s\n", common.JsonMarshal(findAllBook))
}
