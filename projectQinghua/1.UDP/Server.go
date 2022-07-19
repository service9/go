package main

import (
	"fmt"
	"net"
	"os"
)

//处理错误信息
func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func receiveMsg(conn *net.UDPConn) {
	//接收消息
	var buf [30]byte                           //缓冲数组
	n, raddr, err := conn.ReadFromUDP(buf[0:]) //从udp接收数据
	if err != nil {
		//读取错误
		fmt.Println("ERROR", err.Error())
		return
	}
	fmt.Println("MSG IS:", string(buf[0:n]))             //n代表成功读取的字节数
	_, err = conn.WriteToUDP([]byte("hao nimei"), raddr) //回复消息
}
func main1() {
	//创建服务器
	udp_addr, err := net.ResolveUDPAddr("udp", ":8848")
	checkError(err)

	//开启监听
	conn, err := net.ListenUDP("udp", udp_addr)
	defer conn.Close() //关闭连接
	checkError(err)
	receiveMsg(conn) //收消息
}
