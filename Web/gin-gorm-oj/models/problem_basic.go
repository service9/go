package models

import "gorm.io/gorm"

type ProblemBasic struct {
	gorm.Model
	Identity          string             `gorm:"column:identity;type:varchar(36);" json:"identity"`                //问题的唯一标识
	ProblemCategories []*ProblemCategory `gorm:"foreignKey:problem_id;references:id"`                              //关联表
	Title             string             `gorm:"column:title;type:varchar(255);" json:"title"`                     //标题
	Content           string             `gorm:"column:content;type:text;" json:"content"`                         //文章正文
	MaxRuntime        int                `gorm:"column:max_runtime;type:int(11);" json:"max_runtime"`              //最大运行时间
	MaxRunmem         int                `gorm:"column:max_mem;type:int(11);" json:"max_mem"`                      //文章运行内存
	TestCases         []*TestCase        `gorm:"foreignKey:problem_identity;references:identity;" json:"testCase"` //关联表
}

//指定名称
func (table *ProblemBasic) TableName() string {
	return "problem_basic"
}

func GetProblemList(keyword, categoryIdentity string) *gorm.DB {
	tx := DB.Model(new(ProblemBasic)).Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
		Where("title like ? OR content like ?", "%"+keyword+"%", "%"+keyword+"%")
	if categoryIdentity != "" {
		tx.Joins("Right JOIN problem_category pc on pc.problem_id=problem_basic.id").
			Where("pc.category_id=(select cb.id from category_basic cb where cb.identity=?)", categoryIdentity)
	}
	return tx
	// data := make([]*Problem, 0)
	// DB.Find(&data)
	// for _, v := range data {
	// 	fmt.Printf("Problem=>%v\n", v)
	// }
}
