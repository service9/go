package main

import (
	"encoding/base64"
	"fmt"
)

//len(base64Table)=64且所有字节不能重复
const base64Table = "QAZWSXEDCRFVTGBYHNUJMIKOLPqazwsxedcrfvtgbyhnujmikolp1234567890,/"

var coder = base64.NewEncoding(base64Table) //编码解码的表格

func base64Encode(src []byte) []byte { //编码
	return []byte(coder.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) { //解码
	return coder.DecodeString(string(src))
}
func main() {
	fmt.Println(len(base64Table))
	myStr := "我是一男，我爱北大光1"
	deByte := base64Encode([]byte(myStr)) //编码结果
	enByte, err := base64Decode(deByte)   //解码结果
	if err != nil {
		//出错
		fmt.Println("错误:", err)
	}
	fmt.Println(myStr)
	fmt.Println(string(deByte))
	fmt.Println(string(enByte))
}
