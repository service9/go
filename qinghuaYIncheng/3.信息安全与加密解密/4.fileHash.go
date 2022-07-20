package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

//文件加密

func hashSha256File(filePath string) (string, error) {
	var hashValue string //返回hash字符串
	file, err := os.Open(filePath)
	if err != nil {
		//出错
		return hashValue, err //文件打开失败,返回错误,hash为空
	}
	fmt.Println(file)
	defer file.Close()     //延迟关闭文件
	myHash := sha256.New() //创建hash算法对象
	if _, err := io.Copy(myHash, file); err != nil {
		//出错
		return hashValue, err //处理拷贝错误
	}
	hashInBytes := myHash.Sum(nil)
	hashValue = hex.EncodeToString(hashInBytes) //hash转为字符串
	return hashValue, nil
}

func main() {
	filePath := "4文件.txt"
	if hash, err := hashSha256File(filePath); err != nil {
		fmt.Printf("出错了:%s", err)
	} else {
		fmt.Printf("%s sha256hash is %s", filePath, hash)
	}
}
