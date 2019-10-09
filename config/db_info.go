package config

type (
	MongoDBInfo struct {
		URI      string `json:"uri"`
		Database string `json:"database,omitempty" bson:"database"`
	}
	RedisInfo struct {
		Host     string `json:"host,omitempty" bson:"host"`
		Username string `json:"username,omitempty" bson:"username"`
		Password string `json:"password,omitempty" bson:"password"`
	}
)
