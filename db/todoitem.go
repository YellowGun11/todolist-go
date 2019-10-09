package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//go:generate go-easy generate mongodb --type TODOItem --client todoItemCollection
//@def create_at CreatedAt
//@def update_at UpdatedAt
//@def delete_at DeletedAt
type TODOItem struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	OwnerID        primitive.ObjectID `json:"owner_id" bson:"owner_id"`
	TopicID        primitive.ObjectID `json:"topic_id" bson:"topic_id"`
	FatherID       primitive.ObjectID `json:"father_id" bson:"father_id"`
	Content        string             `json:"content" bson:"content"`
	Status         uint32             `json:"status" bson:"status"`
	CompletionTime uint64             `json:"completion_time" bson:"completion_time"`
	CreatedAt      time.Time          `json:"-" bson:"created_at"`
	UpdatedAt      time.Time          `json:"-" bson:"updated_at"`
	DeletedAt      time.Time          `json:"-" bson:"deleted_time"`
	SoftDeleted    bool               `json:"-" bson:"soft_deleted"`
}
