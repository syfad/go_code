package a

import "fmt"

import _ "yuan/day5/MyPro/b"

func func_a()  {
	fmt.Println("a")
}


// 初始化函数
func init()  {
    fmt.Println("a的初始化动作完成!")
}
