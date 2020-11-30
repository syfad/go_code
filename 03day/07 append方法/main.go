package main

import "fmt"

func main() {

	// append方法:  (1) 返回新的切片

	// var numbers []int
	var numbers = []int{1,2,3}

	fmt.Printf("numbers 的cap:%d\n",cap(numbers))
	a := append(numbers, 4)   // [1 2 3 4]
	fmt.Println(a)
	b:=append(a,5)
	fmt.Printf("a 的cap:%d\n",cap(a))
	fmt.Printf("b 的cap:%d\n",cap(b))

	c:=append(b,6,7,8)
	fmt.Printf("c 的cap:%d\n",cap(c))








}
