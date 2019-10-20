package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//go:generate go-easy generate mongodb --type Theme --client todoItemCollection
//@def create_at CreatedAt
//@def update_at UpdatedAt
//@def delete_at DeletedAt
//@def unique_index Owner Name
type Theme struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Owner       primitive.ObjectID `json:"owner" bson:"owner"`
	Name        string             `json:"name" bson:"name"`
	CreatedAt   time.Time          `json:"-" bson:"created_at"`
	UpdatedAt   time.Time          `json:"-" bson:"updated_at"`
	DeletedAt   time.Time          `json:"-" bson:"deleted_time"`
	SoftDeleted bool               `json:"-" bson:"soft_deleted"`
}
