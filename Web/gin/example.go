package main

import "github.com/gin-gonic/gin"

func main2() {
	r := gin.Default() //创建路由引擎
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
