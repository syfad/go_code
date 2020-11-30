package main

import (
	"fmt"
	"net"
	"strings"
)

//保存用户信息的结构体
type Client struct {
	C chan string //传递用户数据
	Name string //用户名（默认与IP地址相同）
	Addr string //客户端的IP
}

var onlineClients = make(map[string]Client) //保存所有用户  {"ip":{"name":"","c":c}}
var message = make(chan string) //传递用户消息

//监听Message通道中的数据
func Manager() {
	for {

		msg := <-message //读取message通道中的数据，如果通道中没有数据，就会一直等待。
		fmt.Println("msg",msg)
		for _, client := range onlineClients {
			fmt.Println("监听:",)
			client.C <- msg

		}
	}
}
func read(conn net.Conn){
	for true{
		data:=make([]byte,1024)
		n, _ :=conn.Read(data)
		fmt.Println(string(data[:n]))
		s:=strings.TrimSpace(string(data[:n]))
		message<-s
	}
}
//处理客户端的连接
func HandleConnect(conn net.Conn) {
	//把客户端的用户信息保存在map对象
	addr := conn.RemoteAddr().String() //获取客户端的IP
	fmt.Println("来自客户端的连接")
	//把用户信息封装成Client
	client := Client{make(chan string), addr, addr}
	onlineClients[addr] = client
	//向所有用户广播消息
	msg := "[" + client.Addr + "]" + client.Name + ": login"
	message <- msg
	//启动WriteMsgToClient的Go程
	go read(conn)
	go WriteMsgToClient(conn, client)
}
//负责向指定的用户发送消息
func WriteMsgToClient(conn net.Conn, client Client) {
	for {
		msg := <- client.C //读取C通道中的数据，如果没有数据，就会一直等待
		conn.Write([]byte(msg + "\n")) //把数据输出到客户端
	}
}

//主协程
func main() {
	fmt.Println("聊天室服务端启动了...")
	//创建一个监听器
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return //结束主协程
	}

	//负责监听Message通道中的数据
	go Manager()

	for {
		conn, err := listener.Accept() //阻塞方法，监听客户端的连接
		if err != nil {
			fmt.Println("listener.Accept err: ", err)
			continue //结束当次循环
		}
		go HandleConnect(conn)
	}

}







