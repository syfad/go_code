package main

import "fmt"

func modify1(x int) {
	x = 100
}

func modify2(p *int) {
	*p = 100
}


func main()  {


	//a:=10
    //b:=&a  // &取地址符号  b此时是一个指针类型的变量,存放的a的地址
	//fmt.Printf("a:%d b:%p\n,type:%T\n", a,b,b) // a:10 b:0xc00000a090 ,type:*int
	//fmt.Println(*b) // *取值操作: *后面是一个指针类型的变量,然后将指针变量的地址对应的值取出来


	a := 10
	modify1(a)
	fmt.Println(a) // 10



	//modify2(&a)
	//fmt.Println(a) // 100





}
