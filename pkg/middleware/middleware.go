package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func MyMid(c *gin.Context){
	fmt.Println("my middleware")
}

func Middleware(c *gin.Context) {
	fmt.Println("this is a middleware!")
}

func CorsHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "123456")
		// c.Next()后就执行真实的路由函数，路由执行完成之后接着走time.Since(t)
		//c.Next()

		// 从time.Now()到目前为止过了多长时间
		latency := time.Since(t)
		log.Print("--", latency)

		// gin设置响应头，设置跨域
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Action, Module, X-PINGOTHER, Content-Type, Content-Disposition")

		//设置中间件的响应头，路由的响应头可以在路由返回中设置，参考/ping
		// c.Writer.WriteHeader(http.StatusMovedPermanently)
		status := c.Writer.Status()
		log.Println("==", status)
	}
}