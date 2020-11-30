package main

import (
	"fmt"
	"net"
)

func main() {
    // 连接服务器
	conn,err:=net.Dial("tcp","127.0.0.1:8899")
	if err != nil {
		fmt.Println("err",err)
	}
	fmt.Println("conn",conn)

	for true {
		// 向服务器发送消息
		buf := make([]byte, 1024)
		fmt.Println(">>>")
		fmt.Scan(&buf)
		conn.Write(buf)     // 网络通信的数据一定是字节串

		// 从服务端接收消息
		res:=make([]byte,1024)
		n,_:=conn.Read(res)
		fmt.Println("data",res)
		fmt.Println("n",n)
		fmt.Println(string(res[:n]))
	}

}



