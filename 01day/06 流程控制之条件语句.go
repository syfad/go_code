package main

import "fmt"

func main() {


    /*

	var score int
    score = 92
    if score >90 {   // 条件语句和 { 必须在一行
      fmt.Printf("成绩优秀!\n")	// 条件为真的时候执行的语句体
      fmt.Printf("nice")	// 条件为真的时候执行的语句体
	}else {
		fmt.Printf("成绩不合格!")
	}

	fmt.Printf("over!")

     */

    // 多分支语法
    /*
	var score int
	score = 77

	if score > 90{
		fmt.Printf("成绩优秀!")
	}else if score >80 {
		fmt.Printf("良好!")
	}else if score >60{
		fmt.Printf("合格!")
	}else{
		fmt.Printf("不合格!")
	}

     */


	// switch语句

	var weeekday = 3

	switch weeekday {
	case 1: fmt.Printf("星期1")
	case 2: fmt.Printf("星期2")
	case 3: fmt.Printf("星期3")
	case 4: fmt.Printf("星期4")
	case 5: fmt.Printf("星期5")
	default:fmt.Printf("周末！")
	}




}