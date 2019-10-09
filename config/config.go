package config

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
)

var config Configure

func init() {
	config.MongoDB = &MongoDBInfo{}
	config.Redis = &RedisInfo{}
	configureFilePath := "./document/configure.json"

	configureByte, err := ioutil.ReadFile(configureFilePath)
	if nil != err {
		logrus.Fatalf("Open configure file %v is err: %v", configureFilePath, err)
	}

	err = json.Unmarshal(configureByte, &config)
	if nil != err {
		logrus.Fatalf("Unmarshal configure is err: %v", err)
	}

	clientOptions := options.Client().ApplyURI(config.MongoDB.URI)
	config.MongoDBClient, err = mongo.NewClient(clientOptions)
	if nil != err {
		logrus.Fatalf("New Mongodb Client is err: %s, URI is : %s", err, config.MongoDB.URI)
	}
	err = config.MongoDBClient.Connect(context.Background())
	if nil != err {
		logrus.Fatalf("connect mongodb is err: %s", err)
	}
}

type Configure struct {
	MongoDB       *MongoDBInfo  `json:"mongodb"`
	Redis         *RedisInfo    `json:"redis"`
	Gin           Gin           `json:"gin"`
	MongoDBClient *mongo.Client `json:"-"`
}

func GetConfigure() Configure {
	return config
}
