package service

import (
	"gin_gorm_oj/define"
	"gin_gorm_oj/helper"
	"gin_gorm_oj/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetRankList
// @Tags     私有方法
// @Summary  分类列表
// @Param    authorization  header    string  true  "token"
// @Param    page           query     int     false  "page"
// @Param    size           query     int     false  "size"
// @Param    keyword        query     string  false  "keyword"
// @Success  200            {string}  json    "{"code":"200","data":""}"
// @Router   /admin/category-list [get]
func GetCategoryList(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetRankList Page Err", err)
		return
	}
	page = (page - 1) * size
	var count int64
	keyword := c.Query("keyword")
	list := make([]*models.CategoryBasic, 0)
	err = models.DB.Model(new(models.CategoryBasic)).
		Where("name like ? ", "%"+keyword+"%").
		Count(&count).Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("GetCategoryList Error:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "GetCategoryList Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}

// CategoryCreate
// @Tags     私有方法
// @Summary  分类创建
// @Param    authorization  header    string  true   "token"
// @Param    name           formData  string  true   "name"
// @Param    parentId       formData  string  false  "parentId"
// @Success  200            {string}  json    "{"code":"200","msg":"","data":""}"
// @Router   /admin/category-create [post]
func CategoryCreate(c *gin.Context) {
	name := c.PostForm("name")
	parentId, _ := strconv.Atoi(c.PostForm("parentId"))
	if name == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数为空",
		})
		return
	}
	category := &models.CategoryBasic{
		Identity: helper.GetUUid(),
		Name:     name,
		ParentId: parentId,
	}
	err := models.DB.Create(&category).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "创建分类失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建分类成功",
	})
}

// CategoryModify
// @Tags     私有方法
// @Summary  分类修改
// @Param    authorization  header    string  true   "token"
// @Param    identity       formData  string  true   "identity"
// @Param    name           formData  string  false  "name"
// @Param    parentId       formData  string  false  "parentId"
// @Success  200            {string}  json    "{"code":"200","msg":"","data":""}"
// @Router   /admin/category-modify [put]
func CategoryModify(c *gin.Context) {
	identity := c.PostForm("identity")
	name := c.PostForm("name")
	parentId, _ := strconv.Atoi(c.PostForm("parentId"))
	if name == "" || identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数为空",
		})
		return
	}
	category := &models.CategoryBasic{
		Identity: identity,
		Name:     name,
		ParentId: parentId,
	}
	err := models.DB.Model(new(models.CategoryBasic)).Where("identity=?", identity).
		Updates(category).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "修改分类失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改分类成功",
	})
}

// CategoryDelete
// @Tags     私有方法
// @Summary  分类删除
// @Param    authorization  header    string  true   "token"
// @Param    identity       query  string  true  "identity"
// @Success  200            {string}  json    "{"code":"200","msg":"","data":""}"
// @Router   /admin/category-delete [delete]
func CategoryDelete(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数为空",
		})
		return
	}
	var count int64
	err := models.DB.Model(new(models.ProblemCategory)).
		Where("category_id=(select id from category_basic where identity=? limit 1)", identity).
		Count(&count).Error
	if err != nil {
		log.Println("ProblemCategory Error:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "获取分类关联的问题失败",
		})
	}
	if count > 0 {
		log.Println("ProblemCategory Error:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "该分类已存在问题，无法删除:" + strconv.FormatInt(count, 10),
		})
		return
	}
	err = models.DB.Where("identity=?", identity).
		Delete(new(models.CategoryBasic)).Error
	if err != nil {
		log.Println("CategoryBasic Delete Error:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "删除分类失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除分类成功",
	})

}
