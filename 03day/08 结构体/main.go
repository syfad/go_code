package main

import "fmt"

type Student struct {
	Id   int
	Name string
	Age  int
	Gender string
}

func main() {
    // 结构体实例化方式1:
    var s1 Student

    fmt.Println(s1)   // 结构体是一个值类型
    // 结构体操作句点符
    s1.Name = "yuan"
    s1.Gender="male"
    fmt.Println(s1.Name)
    fmt.Println(s1.Gender)
    fmt.Println(s1.Id)

    // 实例化方式2:

    s2:=new(Student)
    fmt.Println(s2)
    fmt.Printf("%T",s2)   // 指针类型  *main.Student
    s2.Name = "alex"

	// var x =(*s2).Name
	var x =s2.Name
    fmt.Println(x)

	// 方式3:

	// 结构体作为指针变量实例化
	var s4 *Student = &Student{}
	fmt.Println(s4)
	s4.Name="alvin"
	fmt.Println(s4.Name)
	fmt.Println((*s4).Name)


	// ======================================   结构体的初始化

	// 初始化方式1:

	var s5 =Student{
		Id: 1,
		Name: "yuan",
		Age: 22,
		Gender: "male",
	}


	fmt.Println(s5.Gender)
	fmt.Println(s5.Name)

	var s6=&Student{
		Id: 2,
		Name: "alvin",
		Age: 22,
		Gender: "male",
	}

	fmt.Println(s6.Name)



	// 初始化方式2:多值


	var s7=&Student{3,"yuan",23,"male"}
	fmt.Println(s7.Name)









}
