package main

// 512字节

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	mystr := "sha512 weiweizi QQ1233331"
	mySha512 := sha512.New()
	mySha512.Write([]byte(mystr))
	result := mySha512.Sum(nil)
	fmt.Printf("%x", result)
}
