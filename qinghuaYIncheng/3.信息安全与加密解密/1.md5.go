package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	//数据算法
	myMd5 := md5.New()
	myMd5.Write([]byte("hello baby"))
	result := myMd5.Sum([]byte("")) //不带错误处理就直接加密处理
	fmt.Println(result)
	fmt.Printf("%x\n\n", result) //对比md5的结果实现密码保密、密码校验
}
