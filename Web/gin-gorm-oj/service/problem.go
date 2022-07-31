package service

import (
	"encoding/json"
	"fmt"
	"gin_gorm_oj/define"
	"gin_gorm_oj/helper"
	"gin_gorm_oj/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetProblemList
// @Tags     标签
// @Summary  获取多个标签
// @Param    page               query     int     false  "page"
// @Param    size               query     int     false  "size"
// @param    keyword            query     string  false  "keyword"
// @param    category_identity  query     string  false  "category_identity"
// @Success  200                {string}  json    "{"code":"200","msg":"","data":""}"
// @Router   /problemList [get]
func GetProblemList(c *gin.Context) {
	//获取参数
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
	//page==1 ===> offset 0

	keyword := c.Query("keyword")
	categoryIdentity := c.Query("category_identity")

	//分页
	list := make([]*models.ProblemBasic, 0)
	tx := models.GetProblemList(keyword, categoryIdentity)
	fmt.Println(tx)
	//Omit 忽略列
	err = tx.Count(&count).Omit("content").Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("GetProblemListError", err)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}

// GetProblemDetail
// @Tags     公共方法
// @Summary  问题详情
// @Param    identity  query     string  false  "problem identity"
// @Success  200            {string}  json    "{"code":"200","msg":"","data":""}"
// @Router   /problemDetail [get]
func GetProblemDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": -1,
			"msg":  "问题唯一标识不能为空",
		})
		return
	}
	data := new(models.ProblemBasic)
	err := models.DB.Where("identity=?", identity).
		Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
		First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": -1,
				"msg":  "问题不存在",
			})
		} else {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": -1,
				"msg":  "GetProblemDeatil Err" + err.Error(),
			})
		}
		return
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": -1,
			"data": data,
		})
	}

}

// ProblemCreate
// @Tags     私有方法
// @Summary  问题创建
// @Param    authorization  header    string  true   "token"
// @Param    title          formData  string  true   "title"
// @Param    content        formData  string  true   "content"
// @Param    max_runtime    formData  int     false  "max_runtime"
// @Param    max_memory     formData  int     false  "max_memory"
// @Param    category_ids   formData  array   false  "category_ids"
// @Param    test_cases     formData  array   true   "test_cases"
// @Success  200       {string}  json    "{"code":"200","msg":"","data":""}"
// @Router   /admin/problem-create [post]
func ProblemCreate(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	maxRuntime, _ := strconv.Atoi(c.PostForm("max_runtime"))
	maxMemory, _ := strconv.Atoi(c.PostForm("max_memory"))
	categoryIds := c.PostFormArray("category_ids")
	testCases := c.PostFormArray("test_cases")
	if title == "" || content == "" || len(testCases) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不能为空",
		})
		return
	}
	identity := helper.GetUUid()
	data := &models.ProblemBasic{
		Identity:   identity,
		Title:      title,
		Content:    content,
		MaxRuntime: maxRuntime,
		MaxRunmem:  maxMemory,
	}
	//处理分类
	categoryBasics := make([]*models.ProblemCategory, 0)
	for _, id := range categoryIds {
		categoryId, _ := strconv.Atoi(id)
		categoryBasics = append(categoryBasics, &models.ProblemCategory{
			ProblemId:  data.ID,
			CategoryId: uint(categoryId),
		})
	}
	data.ProblemCategories = categoryBasics
	//处理测试用例
	testCaseBasics := make([]*models.TestCase, 0)
	for _, testCase := range testCases {
		//例子{"input":"1 2\n","output":"3\n"}
		caseMap := make(map[string]string)
		err := json.Unmarshal([]byte(testCase), &caseMap)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "测试用例格式错误" + err.Error(),
			})
			return
		}
		if _, ok := caseMap["input"]; !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "测试用例格式错误 input",
			})
			return
		}
		if _, ok := caseMap["output"]; !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "测试用例格式错误 output",
			})
			return
		}
		testCaseBasic := &models.TestCase{
			Identity:        helper.GetUUid(),
			ProblemIdentity: identity,
			Input:           caseMap["input"],
			Output:          caseMap["output"],
		}
		testCaseBasics = append(testCaseBasics, testCaseBasic)
	}
	data.TestCases = testCaseBasics

	//创建问题
	err := models.DB.Create(data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Create Err:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"identity": data.Identity,
		},
	})
}
