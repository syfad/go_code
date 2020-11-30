package main

import "fmt"

func main() {


	// 变量操作
	var x string
	x = "hello"
	x = "hi"

	fmt.Println(x)
   // 常量赋值 关键字 :const
	const Pi = 3.1415
	// Pi = 123 不能修改常量


	//const (
	//	n1 = 100
	//	n2         // n2 = 100
	//	n3         // n3 = 100
	//)
	//
	//fmt.Println(n1)
	//fmt.Println(n2)
	//fmt.Println(n3)


	// 常量计数器
	//const (
	//	n1 = iota //0
	//	n2        //1
	//	n3        //2
	//	n4        //3
	//)

	const (
		n1 = iota //0
		n2        //1
		_
		n4        //3
	)

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n4)

	// 单行注释

	/*

	多行注释

	 */



}
