package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("网络错误", err.Error())
	}
}

func MessageSend(conn net.Conn) {
	var input string
	for {
		reader := bufio.NewReader(os.Stdin) //读取键盘输入
		data, _, _ := reader.ReadLine()     //读取一行
		input = string(data)                //键盘输入转换为字符串

		if input == "exit" {
			conn.Close()
			fmt.Println("客户端关闭")
			break
		}

		_, err := conn.Write([]byte(input))
		fmt.Println("发送消息成功")
		if err != nil {
			conn.Close()
			fmt.Println("客户端关闭")
			break
		}
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8898") //建立tcp服务器
	defer conn.Close()                             //延迟关闭
	CheckError(err)                                //检查错误
	fmt.Println("连接成功")

	go MessageSend(conn) //开启一个协程

	// conn.Write([]byte("hello nimei"))
	//协程,接收消息
	buf := make([]byte, 1024)
	for {
		numOfBytes, err := conn.Read(buf)
		CheckError(err)
		fmt.Println("收到服务器消息", string(buf[:numOfBytes]))
	}

	fmt.Println("发送消息")
	fmt.Println("客户端关闭")
}
