package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/pkg/dao"
	"test/src/constant"
)

func ProductHandler(c *gin.Context) {
	db := dao.ProductListOrm()
	defer db.Close()

	var resItem [] constant.Menu
	db.Find(&resItem)

	var res constant.ResJson
	res.Header.Code = 200
	res.Header.Msg = "success"
	for _, v := range resItem {
		res.Body = append(res.Body, v)
	}

	//若返回json数据，可以直接使用gin封装好的JSON方法
	c.JSON(http.StatusOK, res)
}

func WelcomeHandler(c *gin.Context) {
	db := dao.ProductListOrm()
	defer db.Close()

	var resItem [] constant.Menu
	db.Find(&resItem)

	var res constant.ResJson
	res.Header.Code = 200
	res.Header.Msg = "success"
	for _, v := range resItem {
		res.Body = append(res.Body, v)
	}

	//若返回json数据，可以直接使用gin封装好的JSON方法
	c.JSON(http.StatusOK, res)
}