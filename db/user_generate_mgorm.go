package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"fmt"
)

var _ = time.Now()
var UserContext = context.Background()

func (model *User) Save() (interface{}, error) {

	model.CreatedAt = time.Now().UTC()

	var count int64
	var err error

	model.ID = primitive.NewObjectID()

	EmailFilter := bson.M{

		"email": model.Email,
	}
	count, err = userCollection.CountDocuments(UserContext, EmailFilter)
	if nil != err {
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf(" Email = %s is exist ",
			model.Email,
		)
	}

	UsernameFilter := bson.M{

		"username": model.Username,
	}
	count, err = userCollection.CountDocuments(UserContext, UsernameFilter)
	if nil != err {
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf(" Username = %s is exist ",
			model.Username,
		)
	}

	result, err := userCollection.InsertOne(UserContext, model)
	if nil != err {
		return nil, err
	}
	return result.InsertedID, err
}

func (model *User) Delete(filter interface{}) (int64, error) {
	result, err := userCollection.DeleteMany(UserContext, filter)
	if nil != err {
		return 0, err
	}
	return result.DeletedCount, err
}

func (model *User) DeleteByID() (int64, error) {
	filter := bson.M{
		"_id": model.ID,
	}
	result, err := userCollection.DeleteOne(UserContext, filter)
	if nil != err {
		return 0, err
	}
	return result.DeletedCount, err
}

func (model *User) FindByID() (err error) {
	filter := bson.M{
		"_id": model.ID,
	}

	result := userCollection.FindOne(UserContext, filter)
	err = result.Decode(model)

	return
}

func (model *User) Find(filter interface{}) (err error) {

	result := userCollection.FindOne(UserContext, filter)
	err = result.Decode(model)
	return
}

func (model *User) FindAll(filter interface{}) (modelList []User, err error) {

	var cursor *mongo.Cursor
	cursor, err = userCollection.Find(UserContext, filter)
	for nil != cursor && cursor.Next(UserContext) {
		temp := User{}
		err = cursor.Decode(&temp)
		if nil != err {
			return
		}
		modelList = append(modelList, temp)
	}
	return
}

func (model *User) FindPage(filter interface{}, iPageSize, iPageIndex int64, SortedStrs ...string) (modelList []User, count int64, err error) {

	count, err = userCollection.CountDocuments(UserContext, filter)
	if nil != err {
		return
	}
	opt := &options.FindOptions{}
	skip := iPageSize * (iPageIndex - 1)
	opt = opt.SetLimit(iPageSize).SetSkip(skip)
	for _, sortStr := range SortedStrs {
		opt = opt.SetSort(sortStr)
	}
	var cursor *mongo.Cursor
	cursor, err = userCollection.Find(UserContext, filter, opt)
	for nil != cursor && cursor.Next(UserContext) {
		temp := User{}
		err = cursor.Decode(&temp)
		if nil != err {
			return
		}
		modelList = append(modelList, temp)
	}
	return
}

func (model *User) UpdateByID() (err error) {
	objectID := model.ID
	filter := bson.M{
		"_id": objectID,
	}

	model.UpdatedAt = time.Now().UTC()

	_, err = userCollection.UpdateOne(UserContext, filter, model)
	return
}

func (model *User) Update(filter interface{}) (err error) {

	model.UpdatedAt = time.Now().UTC()

	_, err = userCollection.UpdateMany(UserContext, filter, model)
	return
}

func (model *User) FindByEmail(email string) (err error) {
	filter := bson.M{

		"email": email,
	}

	result := userCollection.FindOne(UserContext, filter)
	err = result.Decode(model)

	return
}

func (model *User) FindByUsername(username string) (err error) {
	filter := bson.M{

		"username": username,
	}

	result := userCollection.FindOne(UserContext, filter)
	err = result.Decode(model)

	return
}
