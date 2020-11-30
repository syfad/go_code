package main

import "fmt"

type Stu struct {
	name string
	age  int
}

func main()  {
	// 声明管道   引用类型   FIFO
	// var c chan int
    //c:=make(chan int,3)
    //// 插入值
	//c<-12
	//c<-23
	//c<-8


	//fmt.Println(<-c)  // 12
	//fmt.Println(<-c)  // 23
	//fmt.Println(<-c)  // 8
	//<-c  // 取不到值报错

	// 插入任意类型的数据

	c2:=make(chan interface{},10)
	c2<-true
	c2<-"hiyuan"
	c2<-123
	c2<-Stu{"yuan",12}
	//<-c2
	//<-c2
	//<-c2
	//s:=<-c2
    //fmt.Println(s)  // {yuan 12}


    //  遍历管道   之前关闭管道
    // 管道关闭后可读不可写
    // close(c2)
	// fmt.Println(<-c2)

	//for i:=range c2{
	//	fmt.Println(i)
	//}


	// 思考 ?
    // fmt.Println(len(c2))
	for i:=0;i<len(c2);i++ {
		fmt.Println(<-c2)
	}




















}
