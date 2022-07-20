package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

var privateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQC5SL0Mlc2XYMGzuMCYpqQ5Ml/ysJUUZVgFpA0o+M5kID1Efalr
kvAwrACzRm+WSsbEyU/gV3XgsXnaX1JyBUMZqoRJOxOcKETpJwfkviVg+5bDbv8+
iX3hkpiDRPWqirAp5fik8WYT36pJbDccGeCOiIWyxqWzlhxevdn7H2UrUwIDAQAB
AoGBAIkW4pmvNSCt/GPnbMkFczjGpus/7+3ZVhlGWl5YSQ0YNXy5pWLsoz6/5PzA
Psvqo8ryUGFjFNqdNdV29d9bOVs+CkH9mFAFVQCNWTtO4WWhUyd6SXRtrG6tYb71
4f4jhscSAoJ5dhwlOPp5aEKW0drorAA1u4K2zIjHxHy7nSxhAkEAw8LN5z/DpDdE
E8WQmfYTRYZzAdOtoIE9Rx15OW2TmMs+doNJvr3W2y3vwWPUJ0/wQf0UPJwZ6/Lw
NDl8OiHaWQJBAPJMna0fUpl2GAcjfkf/LGpMtOsRvu4n++88bFxYHtslQyeF444t
rvunEAGrD/K13GFLreoSbNR/EQbOaeIC5YsCQQCT59+CVR8QN+Frvt2eNdohsY+7
VZ/doUgXLyGkTjIyQ32SNWfGgdCQEYkQaMimWzN/6CMeGCNNrmPmUXiWw3UhAkEA
7c22PxBr4APghzkVmpHdxKJuOfvbuR/zCpTyARvXhNuociTc9lDt2TzY39pFN7+3
gKJnGUFmIJRpYJW7o7+WCwJAUg8BFtJN4BDSU6OKuy8XHhzgxg5iAIbSfO+y5KMy
NXsGHTsK0qWHDFr4ddhsbLk0xRBZc6IfoFAuQKnJpehwFQ==
-----END RSA PRIVATE KEY-----
`)
var publicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC5SL0Mlc2XYMGzuMCYpqQ5Ml/y
sJUUZVgFpA0o+M5kID1EfalrkvAwrACzRm+WSsbEyU/gV3XgsXnaX1JyBUMZqoRJ
OxOcKETpJwfkviVg+5bDbv8+iX3hkpiDRPWqirAp5fik8WYT36pJbDccGeCOiIWy
xqWzlhxevdn7H2UrUwIDAQAB
-----END PUBLIC KEY-----
`)

func RsaEncrypt(originData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("publicKey is bad")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes) //获取公钥
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, originData) //加密
}

func RsaDecrypt(cipherText []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("privateKey is bad")
	}
	priV, err := x509.ParsePKCS1PrivateKey(block.Bytes) //获取私钥
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priV, cipherText)
}

func main() {
	var decrypted string
	fmt.Println(decrypted) //""
	decrypted = "hello world"
	// flag.StringVar(&decrypted, "d", "s", "描述:加密的数据") //初始化decrypted
	// flag.Parse()
	// fmt.Println(decrypted)
	var data []byte = []byte(decrypted)
	fmt.Println(string(data))
}

func main8() {
	//检查位数
	// var decrypted string
	// flag.StringVar(&decrypted, "d", "", "加密过的数据")
	// flag.Parse()
	// decrypted = "email fdsfsfsd"
	var data []byte
	var err error
	// if decrypted != "" {
	// 	data, err = base64.StdEncoding.DecodeString(decrypted) //提取数据
	// 	if err != nil {
	// 		fmt.Println("错误", err)
	// 	}
	// } else {
	data, err = RsaEncrypt([]byte("QQfsd98fsadjf"))
	if err != nil {
		fmt.Println("错误", err)
	}
	fmt.Println("加密", base64.StdEncoding.EncodeToString(data))
	// }
	originData, err := RsaDecrypt(data) //解密
	if err != nil {
		fmt.Println("错误", err)
	}
	fmt.Println("解密", string(originData))

}
