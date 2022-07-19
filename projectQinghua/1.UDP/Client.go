package main

import (
	"fmt"
	"net"
	"os"
)

func main2() {
	//收发消息 连接网络服务器,ip,端口
	conn, err := net.Dial("udp", "127.0.0.1:8848") //网络类型，地址
	//延时关闭网络
	defer conn.Close()
	if err != nil {
		fmt.Println("net failed")
		os.Exit(1) //code
	}
	//写入发送消息
	conn.Write([]byte("hello nimei"))
	fmt.Println("发送消息", "hello nimei")

	//收取消息
	var msg [30]byte
	conn.Read(msg[0:])

	fmt.Println("received", string(msg[:]))
}
