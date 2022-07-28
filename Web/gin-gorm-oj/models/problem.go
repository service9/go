package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Problem struct {
	gorm.Model
	Identity   string `gorm:"column:identity;type:varchar(36);" json:"identity"`      //问题的唯一标识
	CategoryId string `gorm:"column:categoryId;type:varchar(255);" json:"categoryId"` //分类ID，以逗号分隔
	Title      string `gorm:"column:title;type:varchar(255);" json:"title"`           //标题
	Content    string `gorm:"column:content;type:text;" json:"content"`               //文章正文
	MaxRuntime string `gorm:"column:max_runtime;type:int(11);" json:"max_runtime"`    //最大运行时间
	MaxRunmem  string `gorm:"column:max_mem;type:int(11);" json:"max_mem"`            //文章运行内存
}

//指定名称
func (table *Problem) TableName() string {
	return "problem"
}

func GetProblemList() {
	data := make([]*Problem, 0)
	DB.Find(&data)
	for _, v := range data {
		fmt.Printf("Problem=>%v\n", v)
	}
}
