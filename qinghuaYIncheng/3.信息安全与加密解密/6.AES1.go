package main

import (
	"bytes"
	"fmt"
)

//对称加密 128,不足补充0
//129 32 129=32*4+1
func zeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize //填充的长度
	padText := bytes.Repeat([]byte{'0'}, padding)    //填充0
	return append(ciphertext, padText...)            //字节切片的叠加
}

func zeroUnpadding(originData []byte) []byte { //删除填充
	return bytes.TrimRightFunc(originData, func(r rune) bool {
		return r == '0'
	})
}

func main() {
	myStr := "123456789"
	newBytes := zeroPadding([]byte(myStr), 8)
	fmt.Println(string(newBytes))
	newNewBytes := zeroUnpadding(newBytes)
	fmt.Println(string(newNewBytes))
}
