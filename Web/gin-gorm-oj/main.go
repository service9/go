package main

import (
	"gin_gorm_oj/router"
)

//@title        项目名称
//@version      1.0
//@description  这里是描述
func main() {
	r := router.Router()
	r.Run(":8012")
}
