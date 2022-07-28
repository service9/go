package model

//用户模型 用户结构体
type User struct {
	Id       int
	UserName string
	Password string
	Email    string
}

//用户注册验证模型 用户结构体
type UserReg struct {
	UserName  string `json:"username"`
	Password1 string `json:"Password1"`
	Password2 string `json:"password2"`
	Email     string `json:"email"`
}
