package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type TODOItem struct {
	ID             bson.ObjectId `json:"id" bson:"_id"`
	OwnerID        bson.ObjectId `json:"owner_id" bson:"owner_id"`
	TopicID        bson.ObjectId `json:"topic_id" bson:"topic_id"`
	FatherID       bson.ObjectId `json:"father_id" bson:"father_id"`
	Content        string        `json:"content" bson:"content"`
	Status         uint32        `json:"status" bson:"status"`
	CompletionTime uint64        `json:"completion_time" bson:"completion_time"`
	CreatedAt      time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at" bson:"updated_at"`
	DeletedAt      time.Time     `json:"-" bson:"deleted_time"`
}
