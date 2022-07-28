package admin

import (
	"gin1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//定义结构体
type ArticleController struct {
}

//将方法挂载到结构体上
func (con ArticleController) Add(ctx *gin.Context) {
	//查询数据库
	// userList := []models.User{}
	// models.DB.Find(&userList)

	// 查
	// userList := []models.User{}
	// models.DB.Where("age>2").Find(&userList)

	//增
	// user := models.User{
	// 	Username: "路径",
	// }
	// res := models.DB.Create(&user)

	//改
	// //更新所有字段 推荐
	// user := models.User{ //条件
	// 	Id: 1,
	// }
	// models.DB.Find(&user)
	// user.Username = "哈哈"
	// user.Email = "6988314!@a.ccc"
	// //保存更新
	// models.DB.Save(&user)
	//更新单个列
	// user := models.User{}
	// models.DB.Model(&user).Where("id=?", 1).Update("username", "ssssss啥哈哈")

	//删
	user := models.User{
		Id: 1,
	}
	models.DB.Delete(&user)

	ctx.JSON(http.StatusOK, gin.H{
		// "result": userList,
		// "res": res.Error,
		"user": user,
	})

	ctx.String(200, "我是一个接口")
}
func (con ArticleController) Edit(ctx *gin.Context) {
	ctx.String(200, "我是一个接口")
}
