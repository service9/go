package main

import (
	"fmt"
	"net"
	"strings"
)

//全局变量
//客户端连接，key:ip,端口 value:链接对象
var onlineConns = make(map[string]net.Conn)

//消息队列，缓冲区
var messageQueue = make(chan string, 1000)

//消息，处理程序退出
var quitChan = make(chan bool)

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
		conn.Write([]byte(buf[:numOfBytes]))
		if err != nil {
			break
		}
		if numOfBytes != 0 {
			//消息处理
			message := string(buf[0:numOfBytes])
			//消息队列存储消息
			messageQueue <- message //向通道传数据
		}
	}
}

//消息的协程
func consumeMessage() {
	for {
		select {
		case message := <-messageQueue: //取数据
			doProcessMessage(message) //处理数据
		case <-quitChan: //处理退出
			break
		}
	}
}

//消息的解析函数
func doProcessMessage(message string) {
	contents := strings.Split(message, "#") //字符串切割
	if len(contents) > 1 {
		addr := contents[0]            //取出接收方地址
		sendMessage := contents[1]     //取出消息
		addr = strings.Trim(addr, " ") //地址

		for k, v := range onlineConns {
			var conn = v
			_, err := conn.Write([]byte(sendMessage))
			fmt.Println("服务器发送数据:", k)
			if err != nil {
				fmt.Println("在线链接发送失败:", k)
			}
		}

		//@某人，私聊消息
		// if conn, ok := onlineConns[addr]; ok { //验证连接对象
		// 	_, err := conn.Write([]byte(sendMessage))
		// 	fmt.Println("发送消息", sendMessage, "给地址", addr)
		// 	if err != nil {
		// 		fmt.Println("在线链接发送失败")
		// 	}
		// }
	}
}

func main() {
	//使用协程建立tcp服务器
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8898")
	defer listen_socket.Close() //延迟关闭网络
	CheckError(err)

	fmt.Println("服务器正在等待")

	go consumeMessage()

	for {
		conn, err := listen_socket.Accept() //新的客户端连接
		CheckError(err)
		//处理每一个客户端请求
		addr := fmt.Sprintf("%s", conn.RemoteAddr()) //取出远程地址
		onlineConns[addr] = conn                     //追加连接
		fmt.Println("客户端列表")
		fmt.Println("---------------------------")
		for i := range onlineConns { //循环每一个连接
			fmt.Println(i)
		}

		go ProcessInfo(conn) //发送消息
	}

}
