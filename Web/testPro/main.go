package main

//创建服务器 主要方式

import (
	"fmt"
	"net/http"
	"text/template"
)

//处理器函数 处理请求
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello", r.URL.Path)
}
func testTemplate(w http.ResponseWriter, r *http.Request) {
	//解析模板文件
	// t, _ := template.ParseFiles("template/index.html") //开头不要加/
	//通过Must函数自动处理异常
	t := template.Must(template.ParseFiles("template/index.html",
		"template/index1.html"))
	//上一步相当于
	// t := template.New("路径")
	// t, _ = t.ParseFiles("路径")
	//执行
	// t.Execute(w, "hello") //w
	//指定数据传递的文件
	t.ExecuteTemplate(w, "index1.html", "我要去index1.html中")

}

func main() {
	http.HandleFunc("/", testTemplate) //(根目录,处理器) HandleFunc

	//创建路由
	http.ListenAndServe(":8080", nil) //http1.0
	// http.ListenAndServeTLS() //http2.0 https
}
