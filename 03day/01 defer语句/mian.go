package main

import "fmt"

func f1() int {
	i := 5
	defer func() {
		i++
	}()

	return i
}

func f2() (result int) {
	defer func() {
		result++
	}()
	return 5
}

func f3() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f4() (r int) {
	defer func(r int) {
		r = r + 5
		fmt.Println("r",r)
	}(r)
	return 6
}

func foo(x int)  {
     fmt.Println(x)
}

func main()  {

     // fmt.Printf("start\n")

     // 延迟语句:延迟到return上一步执行,多条defer压栈操作:先进后出
     //defer fmt.Printf("111\n")
	 //defer fmt.Printf("222\n")
	 //defer fmt.Printf("333\n")
     //fmt.Printf("end\n")

	 //fmt.Println(f1())
	 //fmt.Println(f2())
	 //fmt.Println(f3())
	 //fmt.Println("main",f4())

    // 思考题
	//a:=1
    //foo(a)
	//a=2



}
