package config

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/mgorm"
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

	err = mgorm.DefaultMgoInfo(
		config.MongoDB.Hosts,
		config.MongoDB.Database,
		config.MongoDB.Username,
		config.MongoDB.Password,
		config.MongoDB.ConnectTimeoutSecond,
	)
	if nil != err {
		logrus.Fatalf("set mongodb default infomation is err: %v", err)
	}
}

type Configure struct {
	MongoDB *MongoDBInfo `json:"mongodb"`
	Redis   *RedisInfo   `json:"redis"`
	Gin     Gin          `json:"gin"`
}

func GetConfigure() Configure {
	return config
}
