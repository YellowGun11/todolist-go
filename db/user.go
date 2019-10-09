package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//go:generate go-easy generate mongodb --type User --client userCollection
//@def create_at CreatedAt
//@def update_at UpdatedAt
//@def delete_at DeletedAt
//@def unique_index Username
//@def unique_index Email
type User struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Username    string             `json:"username" bson:"username"`
	Password    string             `json:"password" bson:"password"`
	Nickname    string             `json:"nickname" bson:"nickname"`
	Email       string             `json:"email" bson:"email"`
	Header      string             `json:"header" bson:"header"`
	CreatedAt   time.Time          `json:"-" bson:"created_at"`
	UpdatedAt   time.Time          `json:"-" bson:"updated_at"`
	DeletedAt   time.Time          `json:"-" bson:"deleted_time"`
	SoftDeleted bool               `json:"-" bson:"soft_deleted"`
}
