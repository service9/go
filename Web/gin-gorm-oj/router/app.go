package router

import (
	_ "gin_gorm_oj/docs"
	"gin_gorm_oj/middlewares"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gin_gorm_oj/service"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

func Router() *gin.Engine {
	r := gin.Default()
	//配置swagger  url:/swagger/index.html
	// docs.SwaggerInfo.BasePath = "/api/v1"
	// docs.SwaggerInfo.BasePath = "/service"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//路由规则
	//共有方法
	//问题
	r.GET("/problemList", service.GetProblemList)
	r.GET("/problemDetail", service.GetProblemDetail)
	//用户
	r.GET("/userDetail", service.GetUserDetail)
	r.POST("/login", service.Login)
	r.POST("/sendCode", service.SendCode)
	r.POST("/register", service.Register)
	r.GET("/rankList", service.GetRankList) //用户排行榜
	//提交记录
	r.GET("/submitList", service.GetSubmitList)

	//管理员私有方法
	authAdmin := r.Group("/admin", middlewares.AuthAdminCheck())
	//问题创建
	authAdmin.POST("/problem-create", service.ProblemCreate)
	//分类列表
	authAdmin.GET("/category-list", service.GetCategoryList)
	// 分类创建
	authAdmin.POST("/category-create", service.CategoryCreate)
	// 分类修改
	authAdmin.PUT("/category-modify", service.CategoryModify)
	// 分类删除
	authAdmin.DELETE("/category-delete", service.CategoryDelete)
	return r
}
