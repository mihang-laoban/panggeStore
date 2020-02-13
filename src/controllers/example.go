package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHandler(c *gin.Context) {
	//var value string
	//value1, exist1 := c.GetQuery("key")
	//value2, exist2 := c.GetQuery("test")
	//if !exist1 && !exist2 {
	//	value = "the key is not exist!"
	//}
	//if value != "" {
	//	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s \n", value)))
	//}
	//c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s %s\n", value1, value2)))
	type JsonHolder struct {
		Code int `json:"code"`
		Msg string `json:"msg"`
	}
	var resJson JsonHolder
	resJson.Code = 200
	resJson.Msg = "Hello World"
	c.JSON(http.StatusOK, resJson)
}

func PostHandlerTest(c *gin.Context) {
	type JsonHolder struct {
		Id   string    `json:"id"`
		Name string `json:"name"`
	}
	var resJson JsonHolder
	resErr := c.BindJSON(&resJson)
	if resErr != nil {
		panic(0)
	}
	holder := JsonHolder{Id: resJson.Id, Name: resJson.Name}
	//若返回json数据，可以直接使用gin封装好的JSON方法
	c.JSON(http.StatusOK, holder)
}

func PutHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("put success!\n"))
}

func DeleteHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("delete success!\n"))
}