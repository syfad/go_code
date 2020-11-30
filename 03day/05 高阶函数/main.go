package main

import "fmt"

func add_cal(x int,y int){
	fmt.Println(x+y)
}

func foo()  {
	fmt.Printf("foo")
}


func bar(f func()) {
	// f()
	foo()
}


func pay()  {
	fmt.Printf("完成支付功能\n")
}


func ali_pay()  {
	fmt.Printf("ali_pay\n")
}

func wechat_pay()  {
	fmt.Printf("wechat_pay\n")
}


func order_handler(pay func())  {
	pay()
	fmt.Printf("订单结束!")
}


// 函数返回值

func outer () (func()){

	var inner = func (){
		fmt.Printf("OK2")
	}

	return inner
}



func main() {


	//var a = add_cal
	//a(2,3)
	//fmt.Println(a)
	//fmt.Println(add_cal)


	//x:=10
	//y:=10
	//add_cal(x,y)

	// bar(foo)

	// 案例
	// order_handler(ali_pay)
	// order_handler(wechat_pay)

    // 匿名函数的两种调用方式
	//(func (){
	//	fmt.Printf("OK")
	//})()

	// 方式2

	//var f =func (){
	//	fmt.Printf("OK")
	//}
	//f()

	 // go语言要求函数内部声明内部函数必须是匿名函数


	f:=outer()
	f()


}