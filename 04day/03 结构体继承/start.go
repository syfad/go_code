package main

import (
	"fmt"
)

//模拟构造函数
//=============================================================================
type Create struct {
	index string
	times int
}

//问题1：为什么这里是 *create，和直接写create的区别
//构造函数
func create_data(index string, times int) *Create {
	return &Create{
		//index: index,
		//times: times,

		index,times,
	}
}
//方法和接收器
func (c Create) slect() {
	fmt.Printf("%s create done", c.index)
}

//func main() {
//	x := create_data("index_2020", 10)
//	fmt.Println(x.times)
//	x.slect()
//
//}

//==================================================================================

type student struct {
	Name string
	Age int
	class string
	sex string
}

func main() {
	//可简写将*，&去掉
	var s1 *student = &student{Name: "syf",Age: 18, class: "302", sex: "男"}
	fmt.Println(s1,s1.Name,s1.class)
}