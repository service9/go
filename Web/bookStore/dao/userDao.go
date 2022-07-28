package dao

import (
	"bookStore/model"
	"bookStore/utils"
)

//验证登陆 根据用户名密码查询
func CheckLogin(userName string, password string) (*model.User, error) {
	//sql
	sqlStr := "select id,username,password,email from users where username=? and password=?"
	//执行
	row := utils.Db.QueryRow(sqlStr, userName, password) //一条数据
	//扫描结果row
	user := &model.User{}
	row.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
	return user, nil
}

//验证用户是否存在 根据用户名查询
func CheckUserName(userName string) (*model.User, error) {
	//sql
	sqlStr := "select id,username,password,email from users where username=?"
	//执行
	row := utils.Db.QueryRow(sqlStr, userName) //一条数据
	//扫描结果row
	user := &model.User{}
	row.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
	return user, nil
}

//插入用户信息
func Saveuser(userName string, password string, email string) error {
	//sql
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, userName, password, email)
	if err != nil {
		return err
	}
	return nil
}

//验证注册信息
func CheckUserReg(userReg *model.UserReg) bool {
	if len([]rune(userReg.UserName)) < 3 ||
		len(userReg.Password1) < 3 ||
		userReg.Password1 != userReg.Password2 ||
		userReg.Email == "" {
		return false
	}
	return true
}
