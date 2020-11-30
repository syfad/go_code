package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func foo1(c chan int) {
	defer wg.Done()
	fmt.Println("waiting for you")
	x:=<-c
	fmt.Println("x",x)
}


func main() {
	c2:=make(chan int)
	wg.Add(1)
	go foo1(c2)
	time.Sleep(time.Second*3)
	c2<-500
	wg.Wait()
}