package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/go-easy/common/ginx"
	"github.com/siskinc/todolist-go/db"
	"github.com/siskinc/todolist-go/module/errcode"
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
	Email      string `json:"email"`
	Nickname   string `json:"nickname"`
}

func (data *UserRegisterPost) CheckPassword() error {
	if data.Password != data.RePassword {
		return errcode.ResponseErrorCodeRePassword
	}
	return nil
}

func (req *UserRegister) POST(c *gin.Context) {
	data := &UserRegisterPost{}
	err := c.BindJSON(data)
	//err := ginx.BindAll(c, data)

	if nil != err {
		ginx.MakeErrorResponse(c, err)
		return
	}
	err = data.CheckPassword()
	if nil != err {
		ginx.MakeErrorResponse(c, err)
		return
	}
	user := db.User{
		Username: data.Username,
		Password: data.Password,
		Email:    data.Email,
		Nickname: data.Nickname,
	}
	oid, err := user.Save()
	if nil != err {
		ginx.MakeErrorResponse(c, err)
		return
	}
	ginx.MakeResponse(c, oid)
}

type UserLogin struct{}
type UserLoginPost struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserLogin) POST(c *gin.Context) {
	data := &UserLoginPost{}
	err := c.BindJSON(data)
	if nil != err {
		ginx.MakeErrorResponse(c, err)
		return
	}
	logrus.Info(data)
	user := &db.User{}
	err = user.FindByUsername(data.Username)
	if nil != err {
		ginx.MakeErrorResponse(c, err)
		return
	}
	if user.Password != data.Password {
		ginx.MakeErrorResponse(c, err)
		return
	}
	ginx.MakeResponse(c, nil)
	return
}
