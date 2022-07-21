package model

import (
	"Web/utils"
	"fmt"
)

//User结构体
type User struct {
	Id       int
	Username string
	Password string
	Email    string
}

//AddUser 添加User的方法一
func (user *User) AddUser1() error {
	//1、写sql语句
	sqlStr := "insert into users(username,password,email) values(?,?,?);"
	//2、预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译异常:", err)
	}
	//3、执行 并填充占位符 插入不用result
	_, err2 := inStmt.Exec("admin11", "123456", "A@A.A")
	if err2 != nil {
		fmt.Println("执行异常:", err2)
	}
	return nil
}

//AddUser 添加User的方法二
func (user *User) AddUser2() error {
	//1、写sql语句
	sqlStr := "insert into users(username,password,email) values(?,?,?);"
	//不用预编译
	//直接执行
	//3、执行 并填充占位符 插入不用result
	_, err2 := utils.Db.Exec(sqlStr, "admin22", "123456", "A@A.A")
	if err2 != nil {
		fmt.Println("执行异常:", err2)
	}
	return nil
}
