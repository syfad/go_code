package main

import (
	"fmt"
	"sync"
	"time"
)

//  协程
func foo(c chan int){
	defer wg.Done()
    fmt.Println("waiting")
    x:=<-c
    fmt.Println("x",x)

}

var wg sync.WaitGroup

func main()  {
    // 缓冲管道
	// c1:=make(chan int,10)
	// 阻塞通道  无缓冲
	c2:=make(chan int)
    wg.Add(1)
	go foo(c2)
	time.Sleep(time.Second*3)
	c2<-12
    wg.Wait()



}
