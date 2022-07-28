package routers

import (
	"gin1/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(router *gin.Engine) {
	//api路由 组 抽离 分担main.go的压力
	apiRouters := router.Group("/api")
	{
		apiRouters.GET("/index", api.ApiController{}.Index)
		apiRouters.GET("/userList", api.ApiController{}.UserList)
		apiRouters.GET("/pList", api.ApiController{}.PList)
	}
}
