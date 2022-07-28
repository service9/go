package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//定义结构体
type Article struct { //结构体中的属性一定要大写开头，否则返回空结构体
	Title   string `json:"title"` //修改返回数据的key值为小写
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	router := gin.Default() //创建路由
	//配置模板的路径
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(ctx *gin.Context) {
		//路由
		ctx.String(200, "值:%v", "首页")
	})
	router.GET("json", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]interface{}{ //gin.H
			"success": true,
			"msg":     "你好gin",
		}) //返回json数据
	})
	router.GET("json2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": true,
			"msg":     "你好gin",
		}) //返回json数据
	})
	router.GET("json3", func(ctx *gin.Context) {
		a := &Article{ //&取地址，传引用,址传递效率高。反之开辟新空间
			Title:   "我是标题",
			Desc:    "描述",
			Content: "测试内容",
		}
		ctx.JSON(200, a) //返回结构体
	})
	//响应jsonp请求 可以传入回调函数
	//http://localhost:8080/jsonp?callback=xxx
	router.GET("jsonp", func(ctx *gin.Context) {
		a := &Article{ //&取地址，传引用,址传递效率高。反之开辟新空间
			Title:   "我是标题",
			Desc:    "描述",
			Content: "测试内容",
		}
		ctx.JSONP(200, a) //返回jsonp类型的结构体
	})
	//返回xml数据
	router.GET("/xml", func(ctx *gin.Context) {
		a := &Article{ //&取地址，传引用,址传递效率高。反之开辟新空间
			Title:   "我是标题",
			Desc:    "描述",
			Content: "测试内容",
		}
		// ctx.XML(http.StatusOK, gin.H{
		// 	"name":  "name",
		// 	"nam1":  "name",
		// 	"name1": "name",
		// 	"name2": "name",
		// }) //返回xml类型的结构体
		ctx.XML(http.StatusOK, a) //返回xml类型的结构体
	})

	router.GET("/news", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是后台数据",
		})
	})
	router.GET("/goods", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "我是后台数据",
		})
	})
	router.Run() //默认:8080端口
}
