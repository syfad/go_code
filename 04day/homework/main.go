package main

import "fmt"

func dels(s []string)[]string{

	index := 0

	for i:=0;i<len(s);i++{

		if index == 0{

			s[index] = s[i]
            index ++
		}
		if s[index-1] != s[i]{
			s[index] = s[i]
			index++
		}
		fmt.Println("index",index)
	}

	return s[:index]

}





func main() {

	//s:=[]string{"a","b","c","c","d","e","e","e","f"}
	//fmt.Println(dels(s))   // ["a","b","c","d","e"]

	//var a = make([]int,5,10)  // [0,0,0,0,0]
	//fmt.Printf("%p\n",a)
	//b:=a
	//
	//
	//for i:=0;i<10 ;i++  {
	//	a = append(a,i)
	//
	//	if(i==9){
	//		a[0]=1000
	//	}
	//}
	//
	//fmt.Printf("%v\n",a)   //
	//fmt.Printf("%v\n",b)   //
	//fmt.Printf("%p\n",b)   //
	//fmt.Printf("%p\n",a)   //


	//var x = make([]int,5,10)
	//var y = append(x, 100)
	//fmt.Println(x)
	//fmt.Println(y)

}
