package routers

import (
	"gin1/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(router *gin.Engine) { //函数名首字母大写意味着公有
	//admin后台路由 组
	adminRouters := router.Group("/admin")
	{
		adminRouters.GET("/index", admin.IndexController{}.Index)
		adminRouters.GET("/news", admin.IndexController{}.News)
		adminRouters.POST("/user/add", admin.UserController{}.Add)
		adminRouters.POST("/user/edit", admin.UserController{}.Edit)
		adminRouters.POST("/user/edits", admin.UserController{}.Edits)
		adminRouters.POST("/user/uploadImg", admin.UserController{}.UploadImg)
		adminRouters.GET("/article/add", admin.ArticleController{}.Add)
		adminRouters.GET("/article/edit", admin.ArticleController{}.Edit)
	}
}
