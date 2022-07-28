package admin

import "github.com/gin-gonic/gin"

//基本结构体 可被继承 所有的方法
type BaseController struct{}

func (con BaseController) success(ctx *gin.Context) {
	ctx.String(200, "成功")
}
func (con BaseController) error(ctx *gin.Context) {
	ctx.String(400, "失败")
}
