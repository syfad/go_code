package main

import (
	"fmt"
	"sync"
	"time"
)

//
func producer(c chan<- int)  {
	defer wg.Done()
    // c是一个只写的管道
	for i:=0;i<10;i++{
		time.Sleep(time.Second*2)
		c<-i
	}
	close(c)

}


// 从c1管道中读取值平方计算的结果插入c2管道
func consumer(in chan<- int,out <-chan int)  {

    defer wg.Done()
	for val:=range out{
		in<-val*val
	}

	close(in)

}

func print_channel(c <-chan int)  {
	defer wg.Done()
	for val:=range c{
        fmt.Println("c2管道取值:",val)
	}
}


var wg sync.WaitGroup

func main()  {

	// 单向管道
	c1:=make(chan int)
	c2:=make(chan int)

	wg.Add(3)
	go  producer(c1)
	go  consumer(c2,c1)
	go  print_channel(c2)

	wg.Wait()
	fmt.Println("over!")

	



}
