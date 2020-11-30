package b

import "fmt"
import _ "yuan/day5/MyPro/a"

var x=10

// 初始化函数
func init()  {
	fmt.Println("x",x)
	fmt.Println("b的初始化动作完成!")
}
