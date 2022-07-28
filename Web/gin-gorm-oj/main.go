package main

import (
	_ "gin_gorm_oj/docs" // 这里需要引入本地已生成文档
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title        廖圣平博客
// @version      1.0
// @description  swagger 入门使用例子
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/get", Get)
	r.POST("/post", Post)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.GET("/swagger/*any", handleReDoc)
	r.Run() // listen and serve on 0.0.0.0:8080
}

type Response struct {
	Code    uint32      `json:"code"`
	Message uint32      `json:"message" description:"描述"`
	Data    interface{} `json:"data"`
}

type Requests struct {
	Code    uint32      `json:"code"`
	Message uint32      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Code    uint32 `json:"code"`
	Message uint32 `json:"message"`
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.basic  BasicAuth

// @Summary  GET请求的例子
// @Schemes
// @Description  GET请求的例子描述
// @Tags         请求例子
// @Param        id  query  int  true  "Account ID"
// @Accept       json
// @Produce      json
// @Success      200  {object}  Response
// @Router       /get [get]
func Get(c *gin.Context) {
	res := Response{Code: 1001, Message: 1, Data: "connect success !!!"}
	c.JSON(http.StatusOK, res)
}

// @Summary      POST请求的例子
// @Description  POST请求的例子描述
// @Tags         请求例子
// @Param        body  body  Requests  true  "JSON数据"
// @Accept       json
// @Produce      json
// @Success      200  {object}  Response       -->  成功后返回数据结构
// @Failure      400  {object}  ResponseError  -->  失败后返回数据结构
// @Router       /post [post]
func Post(c *gin.Context) {
	res := Response{Code: 1001, Message: 1, Data: "connect success !!!"}
	c.JSON(http.StatusOK, res)
}
