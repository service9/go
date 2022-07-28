package qiantai

import "github.com/gin-gonic/gin"

type DefaultController struct{}

func (con DefaultController) Index(ctx *gin.Context) {
	ctx.String(200, "你是煞笔")
}
func (con DefaultController) User(ctx *gin.Context) {
	ctx.String(200, "你是煞笔")
}
