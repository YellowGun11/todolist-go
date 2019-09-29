package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/todolist-go/config"
)

var Router = gin.New()
var ginConfig = config.GetConfigure().Gin

func init() {
	gin.SetMode(ginConfig.Mode)
}

func Run() {
	err := Router.Run(ginConfig.Addr())

	if nil != err {
		logrus.Fatalf("Server run is err: %v", err)
	}
	logrus.Info("Server is running...")
}
