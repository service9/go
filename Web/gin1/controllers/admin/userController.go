package admin

import (
	"fmt"
	"gin1/models"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//定义结构体 继承
type UserController struct {
	BaseController
}

//将方法挂载到结构体上
func (con UserController) Add(ctx *gin.Context) {
	username := ctx.PostForm("username")
	//单文件
	file, err := ctx.FormFile("avatar")
	if err != nil {
		con.error(ctx)
	} else {
		//保存文件(头像)
		//路径 相对路径都是以main.go为基础
		dst := path.Join("static/upload/avatar", file.Filename)
		err = ctx.SaveUploadedFile(file, dst) //存到static/upload/avatar里
		if err != nil {
			con.error(ctx)
		} else {
			con.success(ctx)
		}
		fmt.Println(username, file.Filename)
	}
}
func (con UserController) Edit(ctx *gin.Context) {
	username := ctx.PostForm("username")
	//多文件上传
	file1, err := ctx.FormFile("avatar1")
	dst1 := path.Join("static/upload/avatar", file1.Filename)
	if err != nil {
		con.error(ctx)
	} else {
		err = ctx.SaveUploadedFile(file1, dst1) //存到static/upload/avatar里
		if err != nil {
			con.error(ctx)
		} else {
			con.success(ctx)
		}
		con.success(ctx)
	}
	file2, err := ctx.FormFile("avatar2")
	dst2 := path.Join("static/upload/avatar", file2.Filename)
	if err != nil {
		con.error(ctx)
	} else {
		err = ctx.SaveUploadedFile(file2, dst2) //存到static/upload/avatar里
		if err != nil {
			con.error(ctx)
		} else {
			con.success(ctx)
		}
		con.success(ctx)
	}
	file3, err := ctx.FormFile("avatar3")
	dst3 := path.Join("static/upload/avatar", file3.Filename)
	if err != nil {
		con.error(ctx)
	} else {
		err = ctx.SaveUploadedFile(file3, dst3) //存到static/upload/avatar里
		if err != nil {
			con.error(ctx)
		} else {
			con.success(ctx)
		}
		con.success(ctx)
	}
	fmt.Println(username, file1.Filename)
	fmt.Println(username, file2.Filename)
	fmt.Println(username, file3.Filename)
	con.success(ctx)
	ctx.JSON(200, gin.H{
		"success":  true,
		"username": username,
		"dst1":     dst1,
		"dst2":     dst2,
		"dst3":     dst3,
	})
}

func (con UserController) Edits(ctx *gin.Context) {
	username := ctx.PostForm("username")
	form, _ := ctx.MultipartForm()
	files := form.File["avatar"]
	for _, file := range files {
		dst := path.Join("static/upload/avatar", file.Filename)
		fmt.Println(file.Filename)
		ctx.SaveUploadedFile(file, dst)
	}
	con.success(ctx)
	fmt.Println(username)
}

func (con UserController) UploadImg(ctx *gin.Context) {
	//按照时间戳命名文件名，安装日期存储到文件夹
	//1获取请求数据
	// username := ctx.PostForm("username")
	file, err := ctx.FormFile("avatar")
	if err != nil {
		con.error(ctx)
	} else {
		//2获取后缀名 筛选指定的出图片格式 .jpg .png .gif .jpeg
		extName := path.Ext(file.Filename)
		allowExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".jpeg": true,
		}
		if _, ok := allowExtMap[extName]; !ok {
			//格式不合法
			//跳转
			ctx.String(200, "上传文件不合法")
			return
		} else {
			//格式合法
			//3创建图片保存的目录 static/upload/20220727
			day := models.GetDay()
			dir := path.Join("static/upload/avatar", day)
			//返回是否创建成功
			err := os.MkdirAll(dir, 0666)
			if err != nil {
				//失败
				ctx.String(200, "MkdirAll失败")
				return
			}
			//4生成文件名 按照时间戳命名
			//获取时间戳
			unix := time.Now().Unix()                         //int64
			fileName := strconv.FormatInt(unix, 10) + extName //转字符串
			//5保存路径和执行保存
			dst := path.Join(dir, fileName)
			ctx.SaveUploadedFile(file, dst)
			con.success(ctx)
		}
	}
}
