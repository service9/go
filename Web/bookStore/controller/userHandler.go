package controller

//路由控制器 处理函数

import (
	"bookStore/dao"
	"bookStore/model"
	"fmt"
	"net/http"
	"text/template"
)

//login处理函数
func Login(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	if username == "" || password == "" {
		t := template.Must(template.ParseFiles("views/pages/login/login.html"))
		t.Execute(w, "")
	}
	//校验登陆
	user, _ := dao.CheckLogin(username, password)
	if user.Id > 0 {
		//用户名和密码正确
		//模板解析 页面跳转
		t := template.Must(template.ParseFiles("views/pages/login/login_success.html"))
		t.Execute(w, "")
	} else {
		t := template.Must(template.ParseFiles("views/pages/login/login.html"))
		t.Execute(w, "")
	}
}

//register处理函数
func Register(w http.ResponseWriter, r *http.Request) {
	//获取post表单信息
	userReg := &model.UserReg{
		UserName:  r.PostFormValue("username"),
		Password1: r.PostFormValue("password1"),
		Password2: r.PostFormValue("password2"),
		Email:     r.PostFormValue("email"),
	}
	if !dao.CheckUserReg(userReg) {
		//验证失败
		fmt.Println("验证失败")
		t := template.Must(template.ParseFiles("views/pages/register/register.html"))
		t.Execute(w, userReg)
		return
	}
	fmt.Println("验证成功")
	user, _ := dao.CheckUserName(userReg.UserName)
	if user.Id > 0 {
		//用户名存在
		t := template.Must(template.ParseFiles("views/pages/register/register.html"))
		t.Execute(w, "")
	} else {
		//用户名可用
		t := template.Must(template.ParseFiles("views/pages/register/register.html"))
		t.Execute(w, userReg)
	}
}
