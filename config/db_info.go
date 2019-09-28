package config

type (
	MongoDBInfo struct {
		Hosts                string `json:"hosts,omitempty" bson:"hosts"`
		ConnectTimeoutSecond int    `json:"connect_timeout_second,omitempty" bson:"connect_timeout_second"`
		Username             string `json:"username,omitempty" bson:"username"`
		Password             string `json:"password,omitempty" bson:"password"`
		Database             string `json:"database,omitempty" bson:"database"`
	}
	RedisInfo struct {
		Host     string `json:"host,omitempty" bson:"host"`
		Username string `json:"username,omitempty" bson:"username"`
		Password string `json:"password,omitempty" bson:"password"`
	}
)
