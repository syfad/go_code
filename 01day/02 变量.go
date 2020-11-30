package main

import "fmt"

func main() {

     // 变量先声明,再使用


	 // 声明一个字符串类型的变量x
	 var x string
	 x = "hello"

     // 一行声明多个变量
     var a,b int
	 a = 1
	 b = 2

	 fmt.Println(x)
	 fmt.Println(a)
	 fmt.Println(b)

	// 声明多个变量another way
	var (
		v5 int
		v6 bool
		v7 string
	)
	 v5 = 5
	 v6 = true
	 v7 = "yuan"
	 fmt.Println(v5)
	 fmt.Println(v6)
	 fmt.Println(v7)



	 //  声明赋值
	 //var v8 = "hello"
	 //v9 := 123

	// 匿名变量
	//var v3,_ = 5,6  // _ 是占位的作用
	//fmt.Println(v3)
	//fmt.Println(_)





}
