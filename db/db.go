package db

import config2 "github.com/siskinc/todolist-go/config"

var (
	config             = config2.GetConfigure()
	database           = config.MongoDB.Database
	todoItemCollection = config.MongoDBClient.Database(database).Collection("todo_item")
	userCollection     = config.MongoDBClient.Database(database).Collection("user")
)
