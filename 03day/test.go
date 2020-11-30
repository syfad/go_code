package main

import "fmt"
//func main()  {
//	var arr = []int{1,2,3,4,5}
//	var s = arr[1:3]
//	s[0] = 10
//	fmt.Println(arr)
//	fmt.Println(arr[1:3])
//}
func main()  {
	var a = make([]int,5,10)
	fmt.Printf("%p\n",a)
	b:=a
	for i:=0;i<10 ;i++  {
		a = append(a,i)

		if i==9{   // i=9
			a[0] =1000
		}

	}
	fmt.Printf("%p\n",a)
	fmt.Printf("%p\n",b)
	fmt.Println(cap(a))
	fmt.Printf("%v\n",a)
	fmt.Printf("%v\n",b)
	fmt.Printf("%p\n",a)
}
