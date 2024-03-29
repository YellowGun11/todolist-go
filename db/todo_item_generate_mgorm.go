// Code generated by go-easy generate-mongodb DO NOT EDIT.
// go-easy url: https://github.com/siskinc/go-easy
package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var _ = time.Now()
var TODOItemContext = context.Background()

func (model *TODOItem) Save() (interface{}, error) {

	model.CreatedAt = time.Now().UTC()

	model.ID = primitive.NewObjectID()

	result, err := todoItemCollection.InsertOne(TODOItemContext, model)
	if nil != err {
		return nil, err
	}
	return result.InsertedID, err
}

func (model *TODOItem) Delete(filter interface{}) (int64, error) {
	result, err := todoItemCollection.DeleteMany(TODOItemContext, filter)
	if nil != err {
		return 0, err
	}
	return result.DeletedCount, err
}

func (model *TODOItem) DeleteByID() (int64, error) {
	filter := bson.M{
		"_id": model.ID,
	}
	result, err := todoItemCollection.DeleteOne(TODOItemContext, filter)
	if nil != err {
		return 0, err
	}
	return result.DeletedCount, err
}

func (model *TODOItem) FindByID() (err error) {
	filter := bson.M{
		"_id": model.ID,
	}

	result := todoItemCollection.FindOne(TODOItemContext, filter)
	err = result.Decode(model)

	return
}

func (model *TODOItem) Find(filter interface{}) (err error) {

	result := todoItemCollection.FindOne(TODOItemContext, filter)
	err = result.Decode(model)
	return
}

func (model *TODOItem) FindAll(filter interface{}) (modelList []TODOItem, err error) {

	var cursor *mongo.Cursor
	cursor, err = todoItemCollection.Find(TODOItemContext, filter)
	for nil != cursor && cursor.Next(TODOItemContext) {
		temp := TODOItem{}
		err = cursor.Decode(&temp)
		if nil != err {
			return
		}
		modelList = append(modelList, temp)
	}
	return
}

func (model *TODOItem) FindPage(filter interface{}, iPageSize, iPageIndex int64, SortedStrs ...string) (modelList []TODOItem, count int64, err error) {

	count, err = todoItemCollection.CountDocuments(TODOItemContext, filter)
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
	cursor, err = todoItemCollection.Find(TODOItemContext, filter, opt)
	for nil != cursor && cursor.Next(TODOItemContext) {
		temp := TODOItem{}
		err = cursor.Decode(&temp)
		if nil != err {
			return
		}
		modelList = append(modelList, temp)
	}
	return
}

func (model *TODOItem) UpdateByID() (err error) {
	objectID := model.ID
	filter := bson.M{
		"_id": objectID,
	}

	model.UpdatedAt = time.Now().UTC()

	_, err = todoItemCollection.UpdateOne(TODOItemContext, filter, model)
	return
}

func (model *TODOItem) Update(filter interface{}) (err error) {

	model.UpdatedAt = time.Now().UTC()

	_, err = todoItemCollection.UpdateMany(TODOItemContext, filter, model)
	return
}
