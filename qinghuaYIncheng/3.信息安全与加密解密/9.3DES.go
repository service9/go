package main

//密钥为24位,3倍的DES

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
)

func DesEncrypt(originData, key []byte) ([]byte, error) {
	//加密
	//修改NewTripleDESCipher
	block, err := des.NewTripleDESCipher(key) //加密密码
	if err != nil {
		return nil, err
	}
	//key修改为前八位(任意八位)
	originData = zeroPadding(originData, block.BlockSize()) //填充位数0
	blockMode := cipher.NewCBCEncrypter(block, key[:8])     //加密模式
	cryped := make([]byte, len(originData))                 //加密存储的空间
	blockMode.CryptBlocks(cryped, originData)               //加密
	return cryped, nil
}

func DesDecrypt(crypted, key []byte) ([]byte, error) {
	//解密
	//修改
	block, err := des.NewTripleDESCipher(key) //解密
	if err != nil {
		return nil, err
	}
	//key修改为前八位(任意八位)
	blockMode := cipher.NewCBCDecrypter(block, key[:8]) //解密模式
	originData := make([]byte, len(crypted))            //解密存储的空间
	blockMode.CryptBlocks(originData, crypted)          //解密
	originData = zeroUnpadding(originData)              //删除尾部补充
	return originData, nil
}

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
	fmt.Println("Game Start")
	key := []byte("123456781234567812345678")           //8个字符
	result, err := DesEncrypt([]byte("123456789"), key) //加密
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result)) //加密结果
	originData, err := DesDecrypt(result, key)             //解密
	if err != nil {
		panic(err)
	}
	fmt.Println(string(originData))

}
