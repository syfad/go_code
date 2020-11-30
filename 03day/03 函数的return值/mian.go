package main

import "fmt"

func  add_cal(x,y int) int{

	//fmt.Println(x+y)
	var z = x+y
	return z
}


func hi(){

	fmt.Printf("hi,yuan!")

}



// 返回多个值
func get_name_age() (string,int) {

	name:= "yuan"
	age:=23

	return name,age
}

// 返回值命名
func  add_cal2(x,y int)(z int){

	//fmt.Println(x+y)
	// z = x+y
	// return
	// return 123
    // return "hi"
	return
}




func main()  {

    //s:=add_cal(12,13)
    //fmt.Println(s)


    //无return情况,仅仅支持调用,不能作为值使用
    //s2:=hi()
    //fmt.Printf(s2)

	//name, _:= get_name_age()
    //fmt.Println(name)


	// 返回值命名
	// fmt.Println(add_cal2(12,23))\



}
