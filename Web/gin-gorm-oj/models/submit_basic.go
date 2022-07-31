package models

import "gorm.io/gorm"

type SubmitBasic struct {
	gorm.Model
	Identity        string        `gorm:"column:identity;type:varchar(36);" json:"identity"`                    //唯一标识
	ProblemIdentity string        `gorm:"column:problem_identity;type:varchar(36);" json:"problem_identity"`    //唯一标识
	ProblemBasic    *ProblemBasic `gorm:"foreignKey:identity;references:problem_identity" json:"problem_basic"` //问题基础表
	UserIdentity    string        `gorm:"column:user_identity;type:varchar(36);" json:"user_identity"`          //唯一标识
	UserBasic       *UserBasic    `gorm:"foreignKey:identity;references:user_identity" json:"user_basic"`       //用户基础表
	Path            string        `gorm:"column:path;type:varchar(255);" json:"path"`                           //代码存放路径
	Status          int           `gorm:"column:status;type:tinyint(1);" json:"status"`
	// 【-1-待判断,1-答案正确,2-错误,3-超时,4-超内存】                     //提交状态
}

func (table *SubmitBasic) TableName() string {
	return "submit_basic"
}

func GetSubmitList(problemIdentity, userIdentity string, status int) (tx *gorm.DB) {
	tx = DB.Model(new(SubmitBasic)).
		Preload("ProblemBasic", func(db *gorm.DB) *gorm.DB {
			return db.Omit("content")
		}).Preload("UserBasic")
	if problemIdentity != "" {
		tx.Where("problem_identity=?", problemIdentity)
	}
	if userIdentity != "" {
		tx.Where("user_identity=?", userIdentity)
	}
	if status != 0 {
		tx.Where("status=?", status)
	}
	return
}
