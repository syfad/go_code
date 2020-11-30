package main

import (
	"fmt"
	"net"
)

type info struct {
	conn net.Conn
	name string
}

var ch_all chan string = make(chan string)
var ch_one chan string = make(chan string)
var ch_who chan string = make(chan string)
var infos map[string]info = make(map[string]info)

func cli(conn net.Conn) {
	for  {
		//从客户端接受消息
		data:=make([]byte,1024)
		n,_:=conn.Read(data)
		//fmt.Println("n", n)
		//fmt.Println("data", data)
		//fmt.Println(string(data[:n]))
		send_info := string(data[:n])
		var oneinfo info
		oneinfo.conn = conn
		oneinfo.name = send_info
		addr := conn.RemoteAddr().String()
		infos[addr] = oneinfo

		ch_all <- send_info



		//发送消息
		//buf:=make([]byte, 1024)
		//fmt.Println("请输入->>>", conn.RemoteAddr())
		//fmt.Scan(&buf)
		//conn.Write(buf)

		msg := <-ch_all
		for _,val :=range infos{
			val.conn.Write([]byte(msg))
		}


	}
}



func main() {
	socket,err := net.Listen("tcp", "0.0.0.0:8800")
	if err != nil{
		fmt.Println("err", err)
	}

	//fmt.Println("conn", conn)
	for true{
		//等待客户端的请求
		fmt.Println("server is waiting...")
		conn,err :=socket.Accept()
		if err != nil{
			fmt.Println("err", err)
		}
		go cli(conn)
	}


}