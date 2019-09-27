package model

import "github.com/globalsign/mgo/bson"

type User struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Username    string        `json:"username" bson:"username"`
	Password    string        `json:"password" bson:"password"`
	Nickname    string        `json:"nickname" bson:"nickname"`
	Email       string        `json:"email" bson:"email"`
	Header      string        `json:"header" bson:"header"`
	IsDeleted   bool          `json:"is_deleted" bson:"is_deleted"`
	DeletedTime uint64        `json:"deleted_time" bson:"deleted_time"`
}
