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