package service

import (
	"gin_gorm_oj/define"
	"gin_gorm_oj/helper"
	"gin_gorm_oj/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUserDetail
// @Tags     公共方法
// @Summary  用户详情
// @Param    identity  query     string  false  "user identity"
// @Success  200    {string}  json    "{"code":"200","msg":"","data":""}"
// @Router   /userDetail [get]
func GetUserDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户唯一标识不能为空",
		})
		return
	}
	data := new(models.UserBasic)
	//省略password字段
	err := models.DB.Omit("password").Where("identity=?", identity).Find(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "GetUserDetail Err" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

// Login
// @Tags     公共方法
// @Summary  用户登陆
// @Param    username  formData  string  false  "username"
// @Param    password  formData  string  false  "password"
// @Success  200       {string}  json    "{"code":"200","msg":"","data":""}"
// @Router   /login [POST]
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填信息为空",
		})
		return
	}
	//md5加密
	password = helper.GetMd5(password)
	print(username, password)
	//校验密码
	data := new(models.UserBasic)
	err := models.DB.Where("name = ? and password = ?", username, password).
		First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或者密码错误",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "Login Err" + err.Error(),
			})
		}
	} else {
		// 生成token
		token, err := helper.GenerateToken(data.Identity, data.Name, data.IsAdmin)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "GenerateToken Error:" + err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"data": gin.H{
					"token": token,
				},
			})
		}
	}
}

// SendCode
// @Tags     公共方法
// @Summary  发送验证码
// @Param    email  formData  string  true  "email"
// @Success  200       {string}  json    "{"code":"200","msg":"","data":""}"
// @Router   /sendCode [POST]
func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "邮箱为空",
		})
	} else {
		//随机生成验证码
		code := helper.GetRand()
		//将验证码保存到redis中
		models.RDB.Set(c, email, code, time.Second*300) //过期时间
		err := helper.SendEmailCode(email, code)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "SendEmailCodeError" + err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "验证码发送成功",
			})
		}
	}
}

// SendCode
// @Tags     公共方法
// @Summary  用户注册
// @Param    mail      formData  string  true   "mail"
// @Param    code      formData  string  true   "code"
// @Param    name      formData  string  true   "name"
// @Param    password  formData  string  true   "password"
// @Param    phone     formData  string  false  "phone"
// @Success  200       {string}  json    "{"code":"200","msg":"","data":""}"
// @Router   /register [POST]
func Register(c *gin.Context) {
	mail := c.PostForm("mail")
	userCode := c.PostForm("code")
	name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	if mail == "" || userCode == "" || name == "" || password == "" || phone == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数为空",
		})
		return
	}
	//验证验证码
	sysCode, err := models.RDB.Get(c, mail).Result()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "出错了,请重新获取验证码",
		})
		return
	}
	if sysCode != userCode {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码不正确",
		})
		return
	}
	//检验邮箱是否存在
	//指定查询的数据表 .Model(new(models.UserBasic))
	var count int64
	err = models.DB.Where("mail=?", mail).Model(new(models.UserBasic)).Count(&count).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "获取用户信息错误" + err.Error(),
		})
		return
	}
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "该邮箱已被注册",
		})
		return
	}
	//数据插入
	identity := helper.GetUUid()
	data := &models.UserBasic{
		Identity: identity,
		Name:     name, //加密
		Password: helper.GetMd5(password),
		Phone:    phone,
		Mail:     mail,
	}
	err = models.DB.Create(data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Create User Error" + err.Error(),
		})
		return
	}
	//生成token
	token, err := helper.GenerateToken(identity, name, data.IsAdmin)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "GenerateToken Error" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}

// GetRankList
// @Tags     公共方法
// @Summary  用户排行榜
// @Param    page  query     int   false  "page"
// @Param    size  query     int   false  "size"
// @Success  200   {string}  json  "{"code":"200","data":""}"
// @Router   /rankList [get]
func GetRankList(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetRankList Page Err", err)
		return
	}
	page = (page - 1) * size
	var count int64

	list := make([]*models.UserBasic, 0)
	err = models.DB.Model(new(models.UserBasic)).Count(&count).Order("finishProblemNum DESC,SubmitNum ASC").
		Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "GetRankList List Err" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": -1,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}
