package main

import "fmt"

func main() {


	 //s:= []int{1,2,3,4,5,6,7}
	 //s1:= s[2:6]
	 //s2:= s[:6]
     //s3:=s[0:5]
     //fmt.Println(s1)   // [3 4 5 6]
     //fmt.Println(s2)   // [1 2 3 4 5 6]
     //fmt.Println(s3)   // [1 2 3 4 5]
	 //
     //fmt.Printf("s1长度%d,s1容量%d",len(s1),cap(s1))  //  s1长度4,s1容量5
     //fmt.Printf("s2长度%d,s2容量%d",len(s2),cap(s2))  //  s2长度6,s2容量7
     //fmt.Printf("s3长度%d,s3容量%d",len(s3),cap(s3))  //  s3长度5,s3容量7


     // 容量练习
	var a=[...]int{1,2,3,4,5,6}
	a1:=a[0:3]  // 3,6
	a2:=a[0:5]  // 5,6
	a3:=a[1:5]  // 4,5
	a4:=a[1:]   // 5,5
	a5:=a[:]    // 6,6
	a6:=a3[1:2] // 1,   4/5
	fmt.Printf("a1的长度%d，容量%d\n",len(a1),cap(a1))
	fmt.Printf("a2的长度%d，容量%d\n",len(a2),cap(a2))
	fmt.Printf("a3的长度%d，容量%d\n",len(a3),cap(a3))
	fmt.Printf("a4的长度%d，容量%d\n",len(a4),cap(a4))
	fmt.Printf("a5的长度%d，容量%d\n",len(a5),cap(a5))
	fmt.Printf("a6的长度%d，容量%d\n",len(a6),cap(a6))

}
