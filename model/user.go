package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Username  string        `json:"username" bson:"username"`
	Password  string        `json:"password" bson:"password"`
	Nickname  string        `json:"nickname" bson:"nickname"`
	Email     string        `json:"email" bson:"email"`
	Header    string        `json:"header" bson:"header"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
	DeletedAt time.Time     `json:"-" bson:"deleted_time"`
}
