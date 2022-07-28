package middlerwares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

//中间件
func InitMiddleware(ctx *gin.Context) {
	//判断用户是否登陆

	fmt.Println(time.Now())
	//请求地址
	fmt.Println(ctx.Request.URL)

	//共享数据
	ctx.Set("username", "张三")

	//定义一个go统计日志
	cCp := ctx.Copy()
	go func() { //协程不影响主线程执行!并发、异步
		time.Sleep(2 * time.Second)
		fmt.Println("Done! in　path" + cCp.Request.URL.Path)
	}()
}
