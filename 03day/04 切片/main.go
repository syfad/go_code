package main

import "fmt"

func main() {
     // 创建切片两种方式


	//  创建切片方式1:  数组切片或切片切片

	//var arr = [5]int{1,2,3,4,5}
	//slice1 := arr[1:3]     // 左闭右开
    //fmt.Println(slice1)  // [2 3]
    //fmt.Printf("%T",slice1)  // []int
	//
    //fmt.Println(arr[0:4]) // []int[1 2 3 4]
    //fmt.Println(arr[0:5]) // []int[1 2 3 4]
    //// fmt.Println(arr[0:6]) // invalid slice index 6 (out of bounds for 5-element array)
	//
    //fmt.Println(arr[2:])   // 缺省 [3 4 5]
    //fmt.Println(arr[:3])   // 缺省 [1 2 3]
    //fmt.Println(arr[:])   // 缺省 [1 2 3 4 5]
	//
    //// fmt.Println(arr[5]) // 最大索引4
    //fmt.Println(arr[1:5])
	//
	//
	//fmt.Println(arr[2:3]) //[3]
	//
	//slice2:= arr[2:5]  // [3,4]
	//fmt.Println(slice2[1]) //  4
	//slice2[1] =10
	//fmt.Println(slice2)
	//fmt.Println(slice1)
	//fmt.Println(arr)
	//
	//slice3:=slice2[0:2]
	//fmt.Println(slice3) // [3 10]

	//  创建切片方式2:  声明定义

	//var slice4 [] int
	//fmt.Println(slice4)   // []
	//fmt.Printf("%T",slice4)  // []int

	// 初始化方式
	var slice5 = []int{1,2,3,4,5}
	fmt.Println(slice5) // [1 2 3 4 5]
	fmt.Printf("%T\n",slice5)  // []int

	//for index, value := range slice5{
	//	fmt.Println(index, value)
	//}

	slice6:= slice5[0:4]
    slice7:= slice6[2:]
    fmt.Println(slice6)  // [1,2,3,4]
    fmt.Println(slice7)  // [3,4]











}
