package admin

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//定义结构体
type IndexController struct {
}

//将方法挂载到结构体上
func (con IndexController) Index(ctx *gin.Context) {
	//session
	//设置sessions
	session := sessions.Default(ctx)
	//options 设置过期时间(单位:秒)
	session.Options(sessions.Options{
		MaxAge: 3600 * 6, //6hours
	})
	session.Set("username", "李四 111")
	session.Save() //必须调用Save保存设置,否则不生效

	//cookie
	//设置cookie(key,vlaue,过期时间(s) -1直接删除,"路径"/,"域名",是否只有https可以访问,可否前端更改)
	// ctx.SetCookie("username", "张三", 10, "/", "localhost", false, false)
	//过期时间如果<0，会立马删除cookie
	// ctx.SetCookie("username", "张三1", -1, "/", "localhost", false, false)
	//多个二级域名共享cookie  .域名
	// ctx.SetCookie("username", "张三", 10, "/", "wei6988314.com", false, false)
	ctx.SetCookie("username", "张三", 10, "/", ".wei6988314.com", false, false)
	ctx.HTML(200, "admin/index.html", "")
	// ctx.String(200, "我是一个接口")
}
func (con IndexController) News(ctx *gin.Context) {
	//获取session
	session := sessions.Default(ctx)
	username1 := session.Get("username")
	ctx.String(200, "我是一个接口,session:%v", username1)

	// //获取cookie
	// username, _ := ctx.Cookie("username")

	// // ctx.HTML(200, "admin/news.html", "")
	// ctx.String(200, "我是一个接口,cookie:"+username)
}
