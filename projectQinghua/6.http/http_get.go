package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main2() {
	url := "http://www.zhangymart.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("错误")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body) //读取信息
	fmt.Println(string(body))
}
