package main

import (
	"bookStore/controller"
	"html/template"
	"net/http"
)

//IndexHandler 去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/index.html"))
	//执行
	t.Execute(w, "")
}

func main() {
	//设置处理静态资源 如js css img
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	//设置处理pages,直接访问静态页面文件
	http.Handle("/pages/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/pages"))))
	http.HandleFunc("/", IndexHandler)
	//登陆
	http.HandleFunc("/login", controller.Login)
	//注册
	http.HandleFunc("/register", controller.Register)
	http.ListenAndServe(":8080", nil)
}
