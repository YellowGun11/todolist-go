package config

import "fmt"

//Gin
type Gin struct {
	Mode               string `json:"mode,omitempty" bson:"mode"`
	Host               string `json:"host,omitempty" bson:"host"`
	Url                string `json:"url,omitempty" bson:"url"`
	Port               int    `json:"port,omitempty" bson:"port"`
	TimeoutReadSecond  int    `json:"timeout_read_second,omitempty" bson:"timeout_read_second"`
	TimeoutWriteSecond int    `json:"timeout_write_second,omitempty" bson:"timeout_write_second"`
}

func (g *Gin) Addr() string {
	return fmt.Sprintf("%s:%d", g.Host, g.Port)
}
