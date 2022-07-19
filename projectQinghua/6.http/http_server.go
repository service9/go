package main

import "net/http"

func main1() {
	//设置服务器响应信息
	//路由
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello http"))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("你没灭啊"))
	})
	//开启服务器
	http.ListenAndServe("127.0.0.1:8080", nil)
}
