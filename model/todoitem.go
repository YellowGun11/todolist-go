package model

import "github.com/globalsign/mgo/bson"

type TODOItem struct {
	ID             bson.ObjectId `json:"id" bson:"_id"`
	OwnerID        bson.ObjectId `json:"owner_id" bson:"owner_id"`
	TopicID        bson.ObjectId `json:"topic_id" bson:"topic_id"`
	FatherID       bson.ObjectId `json:"father_id" bson:"father_id"`
	Content        string        `json:"content" bson:"content"`
	Status         uint32        `json:"status" bson:"status"`
	CompletionTime uint64        `json:"completion_time" bson:"completion_time"`
	DeletedTime    uint64        `json:"deleted_time" bson:"deleted_time"`
}
