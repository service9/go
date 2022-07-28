package router

import (
	_ "gin_gorm_oj/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gin_gorm_oj/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	//配置swagger  url:/swagger/index.html
	// docs.SwaggerInfo.BasePath = "/service"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// r.GET("/", service.Ping)
	r.GET("/problemList", service.GetProblemList)
	return r
}
