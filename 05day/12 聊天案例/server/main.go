package main

import (
	"fmt"
	"net"
)

func handle_client(conn net.Conn)  {
	for true {
		// 从客户端接收消息
		data:=make([]byte,1024)
		n,_:=conn.Read(data)
		fmt.Println("data",data)
		fmt.Println("n",n)
		if(string(data[:n])=="quit"){
			break
		}
		fmt.Println(string(data[:n]))


		// 向客户端响应内容
		res := make([]byte, 1024)
		fmt.Println(">>>")
		fmt.Scan(&res)
		conn.Write(res)     // 网络通信的数据一定是字节串
	}
}

func main() {

	socket,err:=net.Listen("tcp","127.0.0.1:8899")
	if err != nil {
		fmt.Println("err",err)
	}

	// 等待客户端的请求
    fmt.Println("srver is waiting....")

	for true {
		conn,err:=socket.Accept()
		if err != nil {
			fmt.Println("err",err)
		}
		fmt.Println("conn",conn)
		// conn 客户端套接字对象
		go handle_client(conn)
	}



	

}


