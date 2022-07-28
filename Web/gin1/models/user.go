package models

type User struct { //表名 默认操作users表
	//字段
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int //add_time->addTime
}

//配置表名
func (User) TableName() string {
	return "user"
}
