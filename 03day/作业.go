package main

import "fmt"

//第一题 去除相邻重复字符串
func s_remove(slices []string) []string{
	for i := 0; i < len(slices)-1; i++ {
		if slices[i] == slices[i+1] {
			copy(slices[i:],slices[i+1:])
			slices = slices[:len(slices)-1]
		}
	}
	return slices
}

//第二题，旋转函数
func rotat(s []int, i int) []int {
	s = append(s[i:],s[:i]...)
	return s
}


func main()  {
	//去除相邻
	var s = []string{"wo", "wo", "shi", "shi", "syf", "syf", "ABC"}
	x := s_remove(s)
	fmt.Println(x)

	//旋转函数
	var w = []int{1,2,3,4,5}
	fmt.Println(rotat(w, 2))


	//第三题
	var a = make([]int,5,10)
	fmt.Printf("%p\n",a)   //a的指针地址，初始5个0
	b:=a
	for i:=0;i<10 ;i++  {
		a = append(a,i)

		if i==9{   // i=9
			a[0] =1000
		}

	}
	fmt.Printf("%v\n",a) //a被重新赋值
	fmt.Printf("%v\n",b) //切片引用类型地址改变为初始值5个0
	fmt.Printf("%p\n",a) //地址改变
}