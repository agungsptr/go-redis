package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	Id      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title   string             `bson:"title" json:"title"`
	Author  string             `bson:"author" json:"author"`
	Publish BookPublish        `bson:"publish" json:"publish"`
}

type BookPublish struct {
	PublisherName string    `bson:"publisher_name" json:"publisher_name"`
	PublishStatus string    `bson:"publish_status" json:"publish_status"`
	PublishDate   time.Time `bson:"publish_date" json:"publish_date"`
}
