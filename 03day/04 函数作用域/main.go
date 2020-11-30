package main

import "fmt"

// s:=8  // := 不能设置全局变量
var x = 1


func foo()  {
	x:= 10
	fmt.Println("foo",x)
}

func bar()  {
	y:=9
	//fmt.Println("bar",x)
	fmt.Println("bar",y)
}

func main()  {
	x :=100
	foo()
	bar()
	fmt.Println("main",x)

	// if,for是否拥有自己的作用域

	if 2>1 {
		var z =2
		fmt.Println(z)
	}
	// fmt.Println(z)


}
