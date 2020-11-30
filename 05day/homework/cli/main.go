package main

import (
	"fmt"
	"net"
)

func main() {
	conn,err := net.Dial("tcp", "127.0.0.1:8800")
	if err != nil{
		fmt.Println("err", err)
	}

	for  {
		//发送消息
		buf := make([]byte, 1024)
		fmt.Println("请输入->>>")
		fmt.Scan(&buf)
		conn.Write(buf)

		//从服务端接受消息
		res := make([]byte,1024)
		n,_ := conn.Read(res)
		//fmt.Println("data", res)
		//fmt.Println("n", n)
		fmt.Println(string(res[:n]))
	}
}
