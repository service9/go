package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main3() {
	r := gin.Default()                //定义路由引擎
	r.GET("/", func(c *gin.Context) { //配置路由
		c.String(http.StatusOK, "值:%v", "你好") //状态码,格式化string,value
	})
	r.GET("/news", func(c *gin.Context) {
		c.String(http.StatusOK, "值:%v", "我是新闻")
	})
	r.Run(":8000") //启动路由引擎web//默认:8080端口
}
