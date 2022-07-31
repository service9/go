package models

import "github.com/jinzhu/gorm"

//关联表
type ProblemCategory struct {
	gorm.Model
	ProblemId     uint           `gorm:"column:problem_id;type:int(11);" json:"problem_id"`   //问题id
	CategoryId    uint           `gorm:"column:category_id;type:int(11);" json:"category_id"` //分类id
	CategoryBasic *CategoryBasic `gorm:"foreignKey:id;references:category_id"`                //关联表
}

func (table *ProblemCategory) TableName() string {
	return "problem_category"
}
