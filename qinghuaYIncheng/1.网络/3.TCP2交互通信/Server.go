package main

import (
	"fmt"
	"net"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("网络错误", err.Error())
	}
}

//处理请求
func ProcessInfo(conn net.Conn) {
	buf := make([]byte, 1024) //缓存字节数量，开创缓冲区
	defer conn.Close()        //延迟关闭连接
	for {
		numOfBytes, err := conn.Read(buf) //读取数据
		conn.Write([]byte("已收到"))
		if err != nil {
			break
		}
		if numOfBytes != 0 {
			remoteAddr := conn.RemoteAddr()
			fmt.Println("收到消息", string(buf[0:numOfBytes]), "来自:", remoteAddr)
		}
	}
}

func main() {
	//使用协程建立tcp服务器
	listen_socket, err := net.Listen("tcp", "192.168.99.1:8898")
	defer listen_socket.Close() //延迟关闭网络
	CheckError(err)

	fmt.Println("服务器正在等待")

	for {
		conn, err := listen_socket.Accept() //新的客户端连接
		CheckError(err)
		//处理每一个客户端请求
		go ProcessInfo(conn)
	}

}
