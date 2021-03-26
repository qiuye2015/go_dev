package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	//engine := gin.New() //Default调用New
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		context.Writer.Write([]byte("hello world\n"))
	})
	engine.Run()
}
