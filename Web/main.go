package main

//创建服务器 主要方式

import (
	"fmt"
	"net/http"
)

//处理器函数 处理请求
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler) //(根目录,处理器) HandleFunc

	//创建路由
	http.ListenAndServe(":8080", nil) //http1.0
	// http.ListenAndServeTLS() //http2.0 https
}
