package main

import (
	"fmt"
	"net/http"

	"github.com/hyahm/xmux"
)

func main() {
	//创建路由
	router := xmux.NewRouter()
	//忽略斜杠
	// router.IgnoreSlash = true
	//配置路由
	// router.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("<h1>hello xmux</h1>"))
	// })
	// router.Post("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("<h1>hello xmux</h1>"))
	// })

	//合并 get post ...
	router.Get("/{re:([a-z]{1,4})sd: name}", func(w http.ResponseWriter, r *http.Request) {
		name := xmux.Var(r)["name"]
		// age := xmux.Var(r)["age"]
		fmt.Println(name)
		w.Write([]byte(fmt.Sprintf("name:%s", name)))
	})
	// }, http.MethodGet, http.MethodPost)

	router.Run()
}
