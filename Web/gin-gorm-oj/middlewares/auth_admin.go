package middlewares

import (
	"gin_gorm_oj/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 验证是否为管理员身份
func AuthAdminCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO:Check if user is admin
		auth := c.GetHeader("Authorization")
		userClaim, err := helper.AnalyseToken(auth)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized Authorization",
			})
			return
		}
		if userClaim == nil || userClaim.IsAdmin != 1 {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized Admin",
			})
			return
		}
		c.Next()
	}
}
