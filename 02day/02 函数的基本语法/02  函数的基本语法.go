
package main

import "fmt"


// 声明一个函数  func
func hi(){
	fmt.Printf("hi ,yuan!\n")
}


func hi2(name  string){
	fmt.Printf("hi ,%s!\n",name)

}


func hi3(name  string,age int){
	fmt.Printf("姓名:%s,年龄:%d\n",name,age)

}

func  add_cal(x int,y int)  {

	fmt.Println(x+y)
}

func  add_cal2(x,y  int)  {

	fmt.Println(x+y)
}

func  add_cal3(nums ...int)  {

	//fmt.Println("nums:",nums)
	//fmt.Println("len:",len(nums),nums[1])

	s:=0
	for i:=0;i<len(nums);i++ {
		s += nums[i]
	}

	fmt.Println(s)


}




func main()  {

	// 函数调用
	// hi()
	// 再次调用
	// hi()
	//hi2("alvin")
	//add_cal(12,3)   // 函数调用
	//add_cal(1,3)

	// 按位置传参
	// add_cal(3,4,5)   // too many arguments in call to add_cal
	// add_cal(4)  // not enough arguments in call to add_cal

	// hi2()  // go无默认参数

	// 命名实参
	//add_cal(x=3,y=1)
	//add_cal(name="yuan",age=23)

	// hi3(23,"yuan")
	// hi3(name="yuan",age=23)

	// 不定长参数

	add_cal3(12,3,4)
	add_cal3(12,3,4,5)
	add_cal3(12,3)










}


