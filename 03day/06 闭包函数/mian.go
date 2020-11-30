package main

import "fmt"

// 闭包函数:  函数 + 外部变量环境

func adder(z int) func(int) int {
	// sum := 12

	var inner= func(x int) int {
		z += x
		return z
	}
	return inner
}


func main() {
	var i = adder(90)
	fmt.Println(i(78))
}