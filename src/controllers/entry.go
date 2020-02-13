package controllers

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"test/pkg/middleware"
	"test/src/constant"
)

func init(){
	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	if _, err := toml.DecodeFile("./conf/config.toml", &constant.Config); err != nil {
		log.Fatalln(err)
	}
	constant.DbConnection = fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=%s",
		constant.Config.Base.Db["database"],
		constant.Config.Base.Db["password"],
		constant.Config.Base.Db["host"],
		constant.Config.Base.Db["port"],
		constant.Config.Base.Db["username"],
		constant.Config.Base.Db["parseTime"])

	constant.Port = fmt.Sprintf(":%s", constant.Config.Base.Server["port"])
}

func Run() {
	router := gin.Default()    //获得路由实例
	registerMiddleware(router)
	registerHandler(router)

	err := http.ListenAndServe(constant.Port, router)
	if err != nil{
		panic(err)
	}
}

func registerMiddleware(router *gin.Engine)  {
	//添加中间件
	router.Use(middleware.Middleware)
	router.Use(middleware.MyMid)
}

func registerHandler(router *gin.Engine)  {
	//注册接口
	router.GET("/api/home/test", GetHandler)
	router.POST("/api/home/products", ProductHandler)
	router.PUT("/server/put", PutHandler)
	router.DELETE("/server/delete", DeleteHandler)
}