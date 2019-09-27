package config

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
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
		errMsg := fmt.Sprintf("Open configure file %v is err: %v", configureFilePath, err)
		panic(errMsg)
	}
	err = json.Unmarshal(configureByte, &config)
	if nil != err {
		panic(fmt.Sprintf("Unmarshal configure is err: %v", err))
	}
}

type Configure struct {
	MongoDB     *MongoDBInfo         `json:"mongodb"`
	Redis       *RedisInfo           `json:"redis"`
	mongoDBConn *mgorm.MongoDBClient `json:"-"`
	redisConn   *redis.Conn          `json:"-"`
}

func GetConfigure() Configure {
	return config
}
