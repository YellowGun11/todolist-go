package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/siskinc/todolist-go/httpx"
	"net/http"
)

var DemoRouter = Version1Router.Group("/demo/")

func init() {
	httpx.HttpRegister(DemoRouter, "/", &Demo{})
}

type Demo struct {
}

func (demo *Demo) GET(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "here is demo get method",
	})
}

func (demo *Demo) POST(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "here is demo post method",
	})
}

func (demo *Demo) PUT(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "here is demo put method",
	})
}

func (demo *Demo) PATCH(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "here is demo patch method",
	})
}

func (demo *Demo) DELETE(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "here is demo delete method",
	})
}
