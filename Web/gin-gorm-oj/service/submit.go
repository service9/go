package service

import (
	"gin_gorm_oj/define"
	"gin_gorm_oj/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetSubmitList
// @Tags     公共方法
// @Summary  提交列表
// @Param    page             query     int     false  "page"
// @Param    size             query     int     false  "size"
// @param    problemIdentity  query     string  false  "problemIdentity"
// @param    userIdentity     query     string  false  "userIdentity"
// @param    status           query     int     false  "status"
// @Success  200              {string}  json    "{"code":"200","msg":"","data":""}"
// @Router   /submitList [get]
func GetSubmitList(c *gin.Context) {
	size, err := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	if err != nil {
		log.Println("GetProblemList strconv Error", err)
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList strconv Error", err)
	}
	page = (page - 1) * size
	//总条数
	var count int64
	list := make([]models.SubmitBasic, 0)

	problemIdentity := c.Query("problemIdentity")
	userIdentity := c.Query("userIdentity")
	status, _ := strconv.Atoi(c.Query("status"))
	tx := models.GetSubmitList(problemIdentity, userIdentity, status)
	err = tx.Count(&count).Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("GetSubmitList Err:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "GetSubmitList Err:" + err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 200,
			"data": map[string]interface{}{
				"list":  list,
				"count": count,
			},
		})
	}
}
