package models

import "github.com/jinzhu/gorm"

type Submit struct {
	gorm.Model
	Identity        string `gorm:"column:identity;type:varchar(36);" json:"identity"`                 //唯一标识
	ProblemIdentity string `gorm:"column:problem_identity;type:varchar(36);" json:"problem_identity"` //唯一标识
	UserIdentity    string `gorm:"column:user_identity;type:varchar(36);" json:"user_identity"`       //唯一标识
	Path            string `gorm:"column:path;type:varchar(255);" json:"path"`                        //代码存放路径
	Status          string `gorm:"column:status;type:tinyint(1);" json:"Status"`                      //提交状态
}

func (table *Submit) TableName() string {
	return "submit"
}
