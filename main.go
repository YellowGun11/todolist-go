package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/todolist-go/routes/router"
	_ "github.com/siskinc/todolist-go/routes/v1"
)

func init() {
	logrus.SetReportCaller(true)
}

func main() {
	fmt.Println("lalalal")
	logrus.Info("TODOList Server is start...")
	defer logrus.Info("TODOList Server is end.")

	router.Run()
}
