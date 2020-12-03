package main

import (
	"fmt"
	"time"
)

func main()  {


ch3 := make(chan int,3) // 初始化一个带缓冲空间的通道
go readFromChannel(ch3)
go writeToChannel(ch3)

// 主线程休眠1秒，让出执行权限给子 Go 程，即通过 go 开启的 goroutine，不然主程序会直接结束
time.Sleep(1 * time.Second)
}


func writeToChannel(ch chan int) {
	for i := 1; i < 10; i++ {
		fmt.Println("写入：", i)
		ch <- i
	}
}

func readFromChannel(ch chan int) {
	for i := 1; i < 10; i++ {
		v := <-ch
		fmt.Println("读取：", v)
	}
}