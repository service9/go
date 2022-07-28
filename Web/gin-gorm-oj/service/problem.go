package service

import (
	"gin_gorm_oj/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 查询
// @Description 查询
// @Tags 标签
// @Accept 标签
// @Produce 标签
// @Param id path int true "获取多个标签"
// @Success 200 {object} Response
// @Failure 404 {object} Response
// @Router /problemList [get]
func GetProblemList(c *gin.Context) {
	models.GetProblemList()
	c.JSON(http.StatusOK, gin.H{
		"GetProblemList": "GetProblemList",
	})
}
