package main

import "fmt"

func main() {
	//Println与Printf的区别
	//a := 10
	//b := 20
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Printf("abc \n")
	//fmt.Printf("a的值为%d", a)
	//fmt.Println("%d", b) //错误的


	//  { 不能在单独的行上，很有趣的设定
	//var (
	//	name    string
	//	age     int
	//)
	//
	////name = "syf"
	////age = 123
	////Scanf主要分为向外输出内容和获取输入内容两大部分。
	//fmt.Scanf("%s %d", &name, &age)
	//fmt.Printf("扫描结果 name:%s age:%d\n", name, age)

	//变量的多重赋值
	//var (
	//	name string
	//	age int
	//)
	//name, age = "sunyunfeng", 18
	////age = 18
	//fmt.Printf("姓名：%s, 年龄：%d", name, age)

	//常量
	//const pi = 3.1415
	//const e = 2.7182
	//
	//const (
	//	n1 = 100
	//	n2
	//	n3
	//)
	//
	//fmt.Println(n1,n2,pi,e)

	//iota的使用，每增加一行+1
	//const(
	//	n1 = iota
	//	n2 = 100
	//	n3 = iota
	//	n4
	//)
	//const(
	//	a, b = iota + 1, iota +2
	//	c, d
	//	e, f
	//)
	//
	//fmt.Println(n1,n2,n3,n4)
	//fmt.Println(a,b,c,d,e,f)

	//十进制，八进制，十六进制
	//var a int = 10
	//fmt.Printf("%d \n", a)  // 10
	//fmt.Printf("%b \n", a)  // 1010  占位符%b表示二进制
	//
	//// 八进制  以0开头
	//var b int = 077
	//fmt.Printf("%o \n", b)  // 77
	//
	//// 十六进制  以0x开头
	//var c int = 0xff
	//fmt.Printf("%x \n", c)  // ff
	//fmt.Printf("%X \n", c)  // FF

	var hello string
	hello = "hello sunyunfeng"
	 for k, v:= range hello {
	 	fmt.Printf("%d, %c\n", k, v)
	 }
	 
}