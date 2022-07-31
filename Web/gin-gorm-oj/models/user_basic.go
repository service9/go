package models

import "github.com/jinzhu/gorm"

type UserBasic struct {
	gorm.Model
	Identity         string `gorm:"column:identity;type:varchar(36);" json:"identity"`             //用户的唯一标识
	Name             string `gorm:"column:name;type:varchar(100);" json:"name"`                    //用户名
	Password         string `gorm:"column:password;type:varchar(32);" json:"password"`             //密码
	Phone            string `gorm:"column:phone;type:varchar(20);" json:"phone"`                   //手机号
	Mail             string `gorm:"column:mail;type:varchar(100);" json:"mail"`                    //邮箱
	FinishProblemNum int64  `gorm:"column:finishProblemNum;type:int(11);" json:"finishProblemNum"` //完成问题的个数
	SubmitNum        int64  `gorm:"column:submitNum;type:int(11);" json:"submitNum"`               //提交次数
	IsAdmin          int    `gorm:"column:is_admin;type:tinyint(1);" json:"is_admin"`              //是否是管理员 0不是,1是
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
