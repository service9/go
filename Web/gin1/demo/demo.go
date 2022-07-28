package main

import (
	"github.com/gin-gonic/gin"
)

func main1() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
	})
	router.Run() //默认:8080
}
