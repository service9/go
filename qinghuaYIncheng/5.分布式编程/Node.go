package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

//节点数据信息
type NodeInfo struct {
	//节点ID
	NodeId int `json:"nodeId"`
	//节点IP
	NodeIp string `json:"nodeIp"`
	//节点端口
	NodePort string `json:"nodePort"`
}

//节点信息格式化输出
func (node NodeInfo) String() string {
	return "NodeInfo:{nodeId" + strconv.Itoa(node.NodeId) + ",nodeIp:" + node.NodeIp + ",nodePort:" + node.NodePort + "}"
}

//添加节点到集群，请求格式
type AddToClusterMessage struct {
	//源节点
	Source NodeInfo `json:"source"`
	//目的节点
	Dest NodeInfo `json:"dest"`
	//节点连接发送信息
	Message string `json:"message"`
}

//输出信息
func (req AddToClusterMessage) String() string {
	return "AddToClusterMessage:{\n source:" + req.Source.String() + "\n dest:" + req.Dest.String() + ",\n message:" + req.Message + "}"
}

//获取追加到服务器的信息
func getAddToClusterMessage(source NodeInfo, dest NodeInfo, message string) AddToClusterMessage {
	return AddToClusterMessage{
		Source: NodeInfo{
			NodeId:   source.NodeId,
			NodeIp:   source.NodeIp,
			NodePort: source.NodePort,
		},
		Dest: NodeInfo{
			NodeId:   dest.NodeId,
			NodeIp:   dest.NodeIp,
			NodePort: dest.NodePort,
		},
		Message: message,
	}
}

//连接到集群的信息
func connectToCluster(me NodeInfo, dest NodeInfo) bool {
	//连接socket
	connOut, err := net.DialTimeout("tcp", dest.NodeIp+":"+dest.NodePort, time.Duration(10)*time.Second)
	if err != nil {
		if _, ok := err.(net.Error); ok {
			fmt.Println("没有连接到集群", me.NodeId, err)
			return false
		}
	} else {
		fmt.Println("连接到集群，发送消息到节点")
		text := "hi老大，请添加我到节点"
		requestMessage := getAddToClusterMessage(me, dest, text) //请求消息
		json.NewEncoder(connOut).Encode(&requestMessage)         //编码

		decoder := json.NewDecoder(connOut) //解码
		var responseMessage AddToClusterMessage
		decoder.Decode(&responseMessage)
		fmt.Println("得到数据响应\n" + responseMessage.String())
		return true
	}
	return false
}

//接听
func listenOnPort(me NodeInfo) {
	//监听信息
	ln, _ := net.Listen("tcp", fmt.Sprint(":"+me.NodePort))
	for {
		connIn, err := ln.Accept()
		if err != nil {
			if _, ok := err.(net.Error); ok {
				fmt.Println("收消息网络错误", me.NodeId)
			}
		} else {
			var requestMessage AddToClusterMessage
			json.NewDecoder(connIn).Decode(&requestMessage) //解码
			fmt.Println("获得请求\n" + requestMessage.String())
			text := "确定啊小伙子欢迎加入"
			responseMessage := getAddToClusterMessage(me, requestMessage.Source, text) //发送信息给节点
			json.NewEncoder(connIn).Encode(&responseMessage)                           //编码                           //编码
			connIn.Close()                                                             //关闭网络

		}
	}
}

func main() {
	fmt.Println("Game Start...")
	makeMasterOnError := flag.Bool("makeMasterOnError", false, "Ip没有连接集群，我们将本ip设置为master")
	clusterIp := flag.String("clusterIp", "127.0.0.1:8001", "任何节点可以连接")
	myPort := flag.String("myPort", "8001", "正在运行这个节点，端口是8001")
	flag.Parse() //解析命令行参数

	fmt.Println(*makeMasterOnError) //输出信息
	fmt.Println(*clusterIp)
	fmt.Println(myPort)

	//节点生成ID
	rand.Seed(time.Now().UTC().UnixNano()) //随机种子
	myId := rand.Intn(999999999)           //生成随机数
	fmt.Println(myId)

	//抓取ip地址
	myIp, _ := net.InterfaceAddrs()
	fmt.Println(myIp[0].String())

	//创建当前节点信息
	me := NodeInfo{
		NodeId:   myId,
		NodeIp:   myIp[0].String(),
		NodePort: *myPort,
	}
	fmt.Println("当前节点", me)

	dest := NodeInfo{
		NodeId:   -1,
		NodeIp:   strings.Split(*clusterIp, ":")[0],
		NodePort: strings.Split(*clusterIp, ":")[1],
	}
	fmt.Println("目标节点", dest)

	//尝试连接
	ableToConnect := connectToCluster(me, dest)
	//监听其他节点加入集群
	if ableToConnect || (!ableToConnect && *makeMasterOnError) {
		if *makeMasterOnError {
			fmt.Println("当前节点作为主节点")
		}
		listenOnPort(me)
	} else {
		fmt.Println("退出系统，错误")
	}
}
