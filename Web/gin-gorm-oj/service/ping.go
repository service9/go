package service

import "github.com/gin-gonic/gin"

// @Summary 查询
// @Description 查询
// @Tags 标签
// @Accept 标签
// @Produce 标签
// @Param id path int true "获取多个标签"
// @Success 200 {object} Response
// @Failure 404 {object} Response
// @Router /problemList [get]
func Ping(c *gin.Context) { //首字母大写
	c.JSON(200, map[string]interface{}{
		"message": "额",
	})
}
