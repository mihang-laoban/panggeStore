package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/pkg/dao"
	"test/src/constant"
)

func PostHandler(c *gin.Context) {
	ids := dao.ProductList()
	fmt.Println(len(ids))

	type JsonHolder struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Cid int
		Pid int
	}

	type Res struct {
		Code int
		Msg string
		Body [] JsonHolder
	}

	var resJson JsonHolder
	resErr := c.BindJSON(&resJson)
	if resErr != nil {
		panic(0)
	}

	var res Res
	for _, v := range ids{
		holder := JsonHolder{Id: resJson.Id, Name: resJson.Name, Cid: v.Pid, Pid:v.Pid}
		res.Body = append(res.Body, holder)
	}
	res.Code = 200

	//若返回json数据，可以直接使用gin封装好的JSON方法
	c.JSON(http.StatusOK, res)
}


func ProductHandler(c *gin.Context) {
	products := dao.ProductList()

	var res constant.ResJson
	for _, v := range products{
		holder := constant.ItemJson{
			Cid : v.Cid,
			Pid : v.Pid,
			Cname : v.Cname,
			Pname : v.Pname,
			Price : v.Price,
			Energy : v.Energy,
			Img : v.Img,
			Flag : v.Flag,
			Tag : v.Tag,
			Note : v.Note,
		}
		res.Body = append(res.Body, holder)
	}
	res.Header.Code = 200
	res.Header.Msg = "success"

	//若返回json数据，可以直接使用gin封装好的JSON方法
	c.JSON(http.StatusOK, res)
}