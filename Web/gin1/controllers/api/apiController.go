package api

import "github.com/gin-gonic/gin"

type ApiController struct{}

func (con ApiController) Index(ctx *gin.Context) {
	ctx.String(200, "我是一个API接口")
}
func (con ApiController) UserList(ctx *gin.Context) {
	ctx.String(200, "我是一个API接口")
}
func (con ApiController) PList(ctx *gin.Context) {
	ctx.String(200, "我是一个API接口")
}
