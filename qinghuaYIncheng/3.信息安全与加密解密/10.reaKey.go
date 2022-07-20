package main

//实现非对称的加密和解密

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"os"
)

func GenRsaKey(bits int) error {
	//生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	// fmt.Println(privateKey)
	//写入文件
	derStream := x509.MarshalPKCS1PrivateKey(privateKey) //处理privateKey,转成pkcs
	block := &pem.Block{                                 //按照块处理
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	privateFile, err := os.Create("myPrivate.pem") //创建私钥文件
	if err != nil {
		return err
	}
	defer privateFile.Close()            //延时关闭文件
	err = pem.Encode(privateFile, block) //写入文件
	if err != nil {
		return err
	}

	//生成公钥
	publicKey := &privateKey.PublicKey
	fmt.Println(publicKey) //打印公钥
	derpkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derpkix, //模块处理key
	}
	publicFile, err := os.Create("public.pem") //文件流
	if err != nil {
		return err
	}
	defer publicFile.Close()
	err = pem.Encode(publicFile, block) //写入公钥文件
	if err != nil {
		return err
	}

	return nil
}

func main() {
	//生成密钥
	var bits int
	flag.IntVar(&bits, "b", 1024, "密码长度默认1024")
	if err := GenRsaKey(bits); err != nil {
		fmt.Println("密钥文件生成失败")
	}
	fmt.Println("密钥文件生成成功")
}
