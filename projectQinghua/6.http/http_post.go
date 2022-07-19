package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "https://www.baidu.com"
	resp, err := http.Post(url, "application/x-www-form-urlencode",
		strings.NewReader("id=nimei"))
	if err != nil {
		fmt.Println("请求失败")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body) //读取信息
	fmt.Println(string(body))

}
