package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/siskinc/go-easy/common/ginx"
	"github.com/siskinc/todolist-go/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

var ThemeRouter = Version1Router.Group("/theme")

func init() {
	ginx.HttpRegister(ThemeRouter, "", ThemeRest{})
}

type ThemeRest struct{}

type ThemeRestPost struct {
	Name string `json:"name"`
}

type ThemeRestPostResp struct {
	ID interface{} `json:"id"`
}

func (t ThemeRest) POST(c *gin.Context) {
	data := &ThemeRestPost{}
	err := c.BindJSON(data)
	if nil != err {
		ginx.MakeErrorResponse(c, err)
		return
	}
	theme := db.Theme{Name: data.Name}
	insertID, err := theme.Save()
	if nil != err {
		ginx.MakeErrorResponse(c, err)
		return
	}
	ginx.MakeResponse(c, ThemeRestPostResp{ID: insertID})
}

type ThemeRestGet struct {
	QueryIDList []primitive.ObjectID `json:"query_id_list"`
	PageSize    int64                `json:"page_size"`
	PageIndex   int64                `json:"page_index"`
	ThemeName   string               `json:"theme_name"`
}

type ThemeRestGetResp struct {
	ThemeList []db.Theme `json:"theme_list"`
	Count     int64      `json:"count"`
}

func (t ThemeRest) GET(c *gin.Context) {
	data := &ThemeRestGet{}
	err := c.BindJSON(data)
	if nil != err {
		ginx.MakeErrorResponse(c, err)
		return
	}
	theme := &db.Theme{}
	filter := bson.M{}
	if 0 < len(data.QueryIDList) {
		filter["_id"] = bson.E{Key: "$in", Value: data.QueryIDList}
	} else {
		filter["name"] = bson.E{Key: "$regex", Value: data.ThemeName}
	}
	resp := &ThemeRestGetResp{}
	resp.ThemeList, resp.Count, err = theme.FindPage(filter, data.PageSize, data.PageIndex)
	if nil != err {
		ginx.MakeErrorResponse(c, err)
		return
	}
	ginx.MakeResponse(c, resp)
}

type ThemeRestDelete struct {
	ID primitive.ObjectID `json:"id"`
}

func (t ThemeRest) DELETE(c *gin.Context) {
	data := &ThemeRestDelete{}
	err := c.BindJSON(data)
	if nil != err {
		ginx.MakeErrorResponse(c, err)
		return
	}
	theme := &db.Theme{
		ID: data.ID,
	}
	_, err = theme.DeleteByID()
	if nil != err {
		ginx.MakeErrorResponse(c, err)
	}
	c.JSON(http.StatusOK, nil)
}
