package main

import (
	"fmt"
	"github.com/siskinc/todolist-go/config"
)

func main() {
	fmt.Println(config.GetConfigure().MongoDB.Hosts)
}
