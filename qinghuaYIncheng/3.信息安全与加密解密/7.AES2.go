package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// func PKCS5Padding(cipherText []byte, blockSize int) []byte {
// 	padding := blockSize - len(cipherText) //剩余字节
// 	//处理剩余数据
// 	padText := bytes.Repeat([]byte{byte(padding)}, padding)
// 	return append(cipherText, padText...) //追加数据
// }

// func PKCS5Unpadding(originData []byte) []byte {
// 	length := len(originData)
// 	unPadding := int(originData[length-1])   //删除补充
// 	return originData[:(length - unPadding)] //截取
// }

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

//加密
func AesEncrypt(originData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key) //创建密码根据密码加密
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize() //定义大小
	// originData = PKCS5Padding(originData, blockSize) //跨语言
	originData = zeroPadding(originData, blockSize)             //补充0 补充位数
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) //创建加密算法
	crypted := make([]byte, len(originData))                    //加密
	blockMode.CryptBlocks(crypted, originData)                  //按照区块加密
	return crypted, nil
}

//解密
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key) //创建密码
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) //创建解密算法
	originData := make([]byte, len(crypted))                    //创建空间
	blockMode.CryptBlocks(originData, crypted)                  //解密
	originData = zeroUnpadding(originData)                      //删除尾部的0
	return originData, nil
}

func main() {
	fmt.Println("Game Start")
	//AES128  16,24,32 128 192 256
	originData := "QQ7你是猪啊721213"
	key := []byte("1234567890qazwsx")
	result, err := AesEncrypt([]byte(originData), key) //加密
	if err != nil {
		panic(err) //处理异常
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result)) //加密以后的数据
	decryptData, err := AesDecrypt(result, key)            //解密
	if err != nil {
		panic(err) //处理异常
	}
	fmt.Println(string(decryptData)) //解密后的数据
}
