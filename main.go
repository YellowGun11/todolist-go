package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/todolist-go/config"
)

var router = gin.New()
var ginConfig = config.GetConfigure().Gin

func init() {
	logrus.SetReportCaller(true)
	gin.SetMode(ginConfig.Mode)
}

func main() {
	logrus.Info("TODOList Server is start...")
	defer logrus.Info("TODOList Server is end.")

	err := router.Run(ginConfig.Addr())

	if nil != err {
		logrus.Fatalf("Server run is err: %v", err)
	}
}
