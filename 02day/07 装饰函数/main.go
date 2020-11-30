package main

import (
	"fmt"
	"time"
)

// 开放封闭


func foo() int{

	// time_before:=time.Now().Unix()

	time.Sleep(time.Second*3) // 沉睡三秒
	fmt.Printf("foo功能")

	// time_after:=time.Now().Unix()
	// fmt.Println("运行时间：",(time_after - time_before))
    return  11

}


func bar() int{
	time.Sleep(time.Second*5) // 沉睡三秒
	fmt.Printf("bar功能")
	return 22
}


//  方案1
//func timer(f func()){
//
//	time_before:=time.Now().Unix()
//	f()  // 功能函数的执行
//	time_after:=time.Now().Unix()
//	fmt.Println("运行时间：",(time_after - time_before))
//
//}



// 装饰函数

func timer(f func() int) (func() int){
	var inner = func() int{
		time_before:=time.Now().Unix()
		ret:=f() //  功能函数的调用
		time_after:=time.Now().Unix()

		fmt.Println("运行时间：",(time_after - time_before))
		return ret
	}

	return inner

}

func main()  {

    foo :=timer(foo) //   foo :=timer(foo)
    foo()
    bar :=timer(bar)
    bar()

}
