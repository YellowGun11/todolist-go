package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/go-easy/common/ginx"
	"github.com/siskinc/todolist-go/db"
	"net/http"
)

var UserRouter = Version1Router.Group("/user")

func init() {
	ginx.HttpRegister(UserRouter, "/register", &UserRegister{})
	ginx.HttpRegister(UserRouter, "/login", &UserLogin{})
}

type UserRegister struct{}

type UserRegisterPost struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

func (data *UserRegisterPost) CheckPassword() error {
	if data.Password != data.RePassword {
		return fmt.Errorf("两次密码不一致")
	}
	return nil
}

func (req *UserRegister) POST(c *gin.Context) {
	data := &UserRegisterPost{}
	err := ginx.BindAll(c, data)
	if nil != err {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	err = data.CheckPassword()
	if nil != err {
		c.JSON(509, map[string]interface{}{
			"err": err.Error(),
		})
		return
	}
	user := db.User{
		Username: data.Username,
		Password: data.Password,
	}
	oid, err := user.Save()
	if nil != err {
		c.JSON(509, map[string]interface{}{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, oid)
}

type UserLogin struct{}
type UserLoginPost struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AAA      string `form:"aaa"`
}

func (u *UserLogin) POST(c *gin.Context) {
	data := &UserLoginPost{}
	err := ginx.BindAll(c, data)
	if nil != err {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	logrus.Info(data)
	user := &db.User{}
	err = user.FindByUsername(data.Username)
	if nil != err {
		c.JSON(509, map[string]interface{}{
			"err": "用户名或密码错误1",
		})
		return
	}
	if user.Password != data.Password {
		c.JSON(509, map[string]interface{}{
			"err": "用户名或密码错误2",
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "成功",
	})
	return
}
