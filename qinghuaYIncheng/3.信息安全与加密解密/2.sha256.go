package main

//哈希处理字符串
import (
	"crypto/sha256"
	"fmt"
)

func main() {
	mystr := "sha256 weiweizi QQ1233331"
	mySha256 := sha256.New()      //创建加密算法对象
	mySha256.Write([]byte(mystr)) //加入要处理散列的数据
	result := mySha256.Sum(nil)   //计算结果
	fmt.Printf("%x", result)
}
