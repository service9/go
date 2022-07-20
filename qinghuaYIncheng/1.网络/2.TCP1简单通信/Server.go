package TCP

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
		if err != nil {
			break
		}
		if numOfBytes != 0 {
			fmt.Println("收到消息", string(buf))
		}
	}
}

func main() {
	//使用协程建立tcp服务器
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8898")
	defer listen_socket.Close() //延迟关闭网络
	CheckError(err)

	for {
		conn, err := listen_socket.Accept() //新的客户端连接
		CheckError(err)
		//处理每一个客户端请求
		go ProcessInfo(conn)
	}

}
