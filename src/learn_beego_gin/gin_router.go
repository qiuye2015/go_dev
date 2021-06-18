package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main_gin(host string) {
	router := gin.Default()
	// RESTful 路由
	router.GET("/RESTful", helloWorldGet)
	router.POST("/RESTful", helloWorldPost)

	//不支持正则路由
	router.GET("/param/:id", fetchId)

	//组路由
	group1 := router.Group("/g1")
	{
		group1.GET("/action1", action1)
		group1.GET("/action2", action2)
		group1.GET("/action3", action3)
	}
	router.Run(host)
}

// RESTful 路由GET函数
func helloWorldGet(c *gin.Context) {
	c.String(http.StatusOK, "hello, world in GET!")
}

func helloWorldPost(c *gin.Context) {
	c.String(http.StatusOK, "hello, world in POST!")
}

func fetchId(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, fmt.Sprintf("id is :%s\n", id))
}

func action1(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("aciton 1\n"))
}

func action2(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("aciton 2\n"))
}
func action3(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("aciton 3\n"))
}
