package main

import (
    "fmt"
    //"math"
    "strings"
    "strconv"
)

func main() {


	//x := 10
	//fmt.Println("X的值是%d.",x) // X的值是%d. 10
	//fmt.Printf("X的值是%d.",x) // X的值是%d. 10
	//fmt.Printf("X的值是%d.",x) // Printf可以形成占位打印但是没有默认换行

	//fmt.Printf("X的十进制值是%d.",x) // 10
	//fmt.Printf("X的二进制值是%b.",x) // 1010
	//fmt.Printf("X的八进制值是%o.",x) // 12
	//fmt.Printf("X的十六进制值是%x.",x) // a
	//
	//
	//fmt.Printf("%f\n", math.Pi)
	//fmt.Printf("%.5f\n", math.Pi)


	// 布尔类型

	// true
	// false

	//b1:= (2>3) || (3==3)
	//fmt.Println(b1)
	//
	//b2:= !(2>3) || (3>5)
	//fmt.Println(b2)
	//fmt.Printf("b2的结果类型%T",b2)

    // 字符串类型

    //s := "yuan"   //   字符串一定是双引
    //fmt.Println(s)
    //s1 := "hi "
    //s2 := "yuan!"
    //s3 := s1 + s2
    //fmt.Println(s3)
    //
    //fmt.Println("str := \"c:\\Code\\lesson01\\go.exe\"")
    fmt.Printf("hi\n")
    //fmt.Printf("yuan")
    // 多行打印: 反引号


    s4 :=`请选择商品序号:
                    1  苹果手机\n
                    2  mac
                    3  冰箱`

    fmt.Println(s4)

    //  字符串的常用方法
    s5 := "hi yuan!"
    size := len(s5)
    fmt.Println(size)

    s6 := "zhangsan-li-wangwu"
    s7 := strings.Split(s6,"-") // [zhangsan li wngwu]
    fmt.Println(s7)
    fmt.Println(len(s7))

    s8 := strings.Join(s7," ")
    fmt.Println(s8)


    b := strings.Contains(s6,"wangwu")
    fmt.Println(b)
    b1 := strings.HasPrefix(s6,"z")
    fmt.Println(b1)
    i1 := strings.Index(s6,"an")
    fmt.Println(i1)


    // 类型转换
    f1 := float64(12)
    fmt.Println(f1)

    s9 := strconv.Itoa(123) + "yuan"   // 将int类型转换为string
    fmt.Println(s9)
    i2,err := strconv.Atoi("1200")    // 将字符串类型转换为整形
    fmt.Println(i2)
    fmt.Println(err)





















}

