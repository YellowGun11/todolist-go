package main

import (
	"github.com/sirupsen/logrus"
	"github.com/siskinc/todolist-go/routes/router"
)

func init() {
	logrus.SetReportCaller(true)
}

func main() {
	logrus.Info("TODOList Server is start...")
	defer logrus.Info("TODOList Server is end.")

	router.Run()
}
