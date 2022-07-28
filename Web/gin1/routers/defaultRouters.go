package routers

import (
	"fmt"
	"gin1/controllers/qiantai"
	"gin1/middlerwares"
	"time"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(router *gin.Engine) {
	//前台路由 组
	defaultRouters := router.Group("/default")
	//中间件法二、
	defaultRouters.Use(middlerwares.InitMiddleware)
	{
		//中间件 定义在处理函数前后
		defaultRouters.GET("/index", func(ctx *gin.Context) {
			//计算请求处理时间
			start := time.Now().UnixNano()
			ctx.Abort() //终止执行该请求其他剩余函数(不包括本函数)
			ctx.Next()  //执行该请求其他剩余函数
			end := time.Now().UnixNano()
			username, _ := ctx.Get("username") //返回空接口类型,使用需要转类型
			v, ok := username.(string)         //将空接口类型转为string类型
			if ok == true {
				ctx.String(200, "获取成功:"+v)
			} else {
				ctx.String(200, "获取失败")
			}
			fmt.Println("username:", username)
			fmt.Println("我是中间件1", end-start) //最后执行了
		}, func(ctx *gin.Context) {
			fmt.Println("我是中间件2")
		}, qiantai.DefaultController{}.Index, func(ctx *gin.Context) {
			fmt.Println("我是中间件3")
		})
		defaultRouters.GET("/user", qiantai.DefaultController{}.User)
	}
}
