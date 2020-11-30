package main

import "fmt"

//
//func timer(f func(x,y int)) (func(int,int)) {
//
//	wrapper := func(x,y int) {
//		time_before:=time.Now().Unix()
//		f(x,y)
//		time_after:=time.Now().Unix()
//		fmt.Println("运行时间：",(time_after - time_before))
//	}
//
//	return wrapper
//
//}
//
//func add(x,y int)  {
//	fmt.Printf("%d+%d=%d",x,y,x+y)
//	time.Sleep(time.Second*2)
//}
//
//func mul(x,y int)  {
//	fmt.Printf("%d*%d=%d",x,y,x*y)
//	time.Sleep(time.Second*3)
//}
//
//func main() {
//
//	var add = timer(add)
//	add(1,4)
//	var mul =timer(mul)
//	mul(2,5)
//
//}



//var a,b = 0,1
//
//func foo()  {
//	for i := 1; i<15;i++{
//		a,b = b,a+b
//
//		fmt.Println(a)
//	}
//
//}
//
//
//
//func f(num int) int {
//	if num == 1 || num == 2 {
//		return 1
//	}
//	return f(num-1) + f(num-2)
//}
//
//func main()  {
//	fmt.Println(f(5))
//	foo()
//}

//func pay(){
//	fmt.Printf("pay func")
//}
//
//func order_handler(f func()) {
//	fmt.Printf("功能")
//	pay()
//}
//
//func main()  {
//	order_handler(pay)
//}

func add(x,y int) int{
	return x+y
}

func mul(x,y int) int{
	return x*y
}

func cal(x,y int,f func(x,y int) int) int{
	return f(x,y)
}

func main() {
	ret1 := cal(12,3,add)
	fmt.Println(ret1)
	ret2 := cal(12,3,mul)
	fmt.Println(ret2)
}

