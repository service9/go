package main

import (
	"gin1/models"
	"gin1/routers"
	"html/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"

	"github.com/gin-gonic/gin"
)

//Session 保存在服务器中
//key返回到客户端 value保存在服务器
//key写入Cookie中
//http是无状态的
//导入第三方中间件实现
//如果是多台服务器(分布式部署)，需要将session保存在数据库中redis,mysql
//比如：将cookie引擎换成redis引擎

func main() {
	//创建路由引擎
	router := gin.Default() //推荐额外加载日志和错误
	// router := gin.New()
	//设置内存限制，默认32MiB
	router.MaxMultipartMemory = 8 << 20 //8MiB
	//自定义模板函数 在加载模板前
	router.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})
	//加载模板文件
	router.LoadHTMLGlob("templates/**/*")
	//配置静态web目录
	router.Static("/static", "static")
	//全局中间件 优先于局部中间件执行
	// router.Use(func(ctx *gin.Context) {
	// 	fmt.Println("你是猪")
	// })
	//创建cookie存储引擎，可以替换为其他存储引擎
	// store := cookie.NewStore([]byte("secret999")) //密钥
	//分布式部署使用redis存储引擎
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret999"))
	// 全局注册session中间件
	router.Use(sessions.Sessions("mySession", store))

	//导入路由组的初始化函数
	routers.DefaultRoutersInit(router)
	routers.AdminRoutersInit(router)
	routers.ApiRoutersInit(router)
	router.Run(":80")
}
