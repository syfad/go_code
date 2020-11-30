package main

import "fmt"

//func main()  {
//	fmt.Printf("hello sun")
//}

func main() {
	//
	//// 方式1:先声明再赋值
	//var name string
	//name = "yuan"
	//// 方式2: 声明并赋值
	//var age = 18
	//
	//// 声明并赋值简写
	//gender := "male"
	//
	//// 多重赋值
	//var v1,v2 int
	//v1,v2 = 1,2
	//
	//// 匿名变量
	//var v3,_ = 5,6  // _ 是占位的作用
	//
	//// 不能重复声明变量
	//// var v3 = 100
	//
	//fmt.Println(name)
	//fmt.Println(age)
	//fmt.Println(gender)
	//fmt.Println(v1)
	//fmt.Println(v2)
	//fmt.Println(v3)


	// 0 1 1 2 3 5 8 13
	a,b :=0,1
	sum1:=0
	for i:=0;i<60;i++{
		a,b =b, a+b
		// fmt.Println(a)
		sum1 = a+b
		//a = b
		//b = sum1
		fmt.Println(sum1)

	}




	//k1 == 0
	//for i :=1;i<61;i++{
	//	k1,k2 == i, i-1
	//
	//	//if i < 2{
	//	//	fmt.Println(i)
	//	//}
	//
	//	fmt.Println(i)
	//}



	//1 1 2 3 5 8 13 21

}