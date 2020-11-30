package main

import "fmt"

func main()  {
	//var c chan int
	//c = make(chan int, 3)
	//c<-10
	//c<-5
	//c<-90
	//fmt.Println(<-c)
	//close(c)
	//
	//for i:=range c{
	//	fmt.Println(i)
	//}

	ch := make(chan int)
	ch <- 20
	fmt.Println("success")
}