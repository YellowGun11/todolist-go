package model

import "github.com/globalsign/mgo/bson"

type TODOItem struct {
	ID             bson.ObjectId `json:"id" bson:"_id"`
	OwnerID        bson.ObjectId `json:"owner_id" bson:"owner_id"`
	TopicID        bson.ObjectId
	FatherID       bson.ObjectId
	Content        string
	Status         uint32
	CompletionTime uint64
	DeletedTime    uint64
}
