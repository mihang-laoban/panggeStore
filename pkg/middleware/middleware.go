package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func MyMid(c *gin.Context){
	fmt.Println("my middleware")
}

func Middleware(c *gin.Context) {
	fmt.Println("this is a middleware!")
}

// 中间件
func CorsHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		// gin设置响应头，设置跨域
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Action, Module, X-PINGOTHER, Content-Type, Content-Disposition")
	}
}