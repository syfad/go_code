package main

import "fmt"

func foo(){

	fmt.Printf("foo")
	// 自己调用自己
	foo()
	fmt.Printf("end")

}

func f(n int) int{
    if n == 1{
    	return 1
	}
	// n*(n-1)*(n-2)....*1
	return n*f(n-1)

}



//  1 1 2 3 5 8 13  21
//  1 2 3 4 5 6 7   8

func fib(n int) int{

    if n==1 || n==2{
    	return 1
	}

	return fib(n-1)+fib(n-2)
}


func main(){

     // 实现n的阶乘
	// fmt.Println(f(10))

	fmt.Println(fib(8))


}
