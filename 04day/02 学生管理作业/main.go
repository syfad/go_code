package main

import (
	"fmt"
	"os"
)

//



type student struct {
	sid    int
	name   string
	class  string
	score  map[string] int
}

func newStudent(sid int,name string,class  string,score map[string] int)(*student){
	return &student{
		sid:   sid,
		name:  name,
		class: class,
		score: score,
	}
}

//var students [] *student


//  设计学生管理结构体   ===============================================================

type StuManager struct {
	students_list [] *student
}

func (s *StuManager) add_stu()  {
	fmt.Println("添加学生成绩！")
	new_stu:=get_input_stu()
	s.students_list=append(s.students_list,new_stu)

}

func (s *StuManager) select_all_stus()  {
	fmt.Println("查询所有学生成绩！")

	for _,stu :=range s.students_list{
		fmt.Printf("学号：%5d，姓名：%-5s，班级，%s，语文成绩：%5d，数学成绩%5d，英文成绩%-5d，平均成绩%5d\n",stu.sid,stu.name,stu.class,stu.score["chinese"],stu.score["math"],stu.score["english"],stu.score["average"])

	}

}

func  (s *StuManager) edit_stu()  {
	new_stu:=get_input_stu()

	for i:=0;i<len(s.students_list);i++{
		if  s.students_list[i].sid ==new_stu.sid{
			s.students_list[i] = new_stu
		}
	}

	fmt.Println("编辑学生成绩！")
}

func (s *StuManager) score_sort()  {
	fmt.Println("排序学生成绩！")

	vLen := len(s.students_list)
	for i := 0; i < vLen-1; i++ {
		for j := 0; j < vLen-i-1; j++ {
			if s.students_list[j].score["average"] > s.students_list[j+1].score["average"] {
				s.students_list[j], s.students_list[j+1] = s.students_list[j+1], s.students_list[j]
			}
		}
	}

	s.select_all_stus()

}


// ===============================================================








func show_menu()  {
	fmt.Println("欢迎来到学生成绩管理系统！")
	fmt.Println("1、添加学生成绩")
	fmt.Println("2、查询所有学生成绩")
	fmt.Println("3、编辑学生成绩")
	fmt.Println("4、成绩从高到低排序！")
	fmt.Println("5、退出程序！")

}

func get_input_stu()(*student){

	var sid,chinese,math,english int
	var name string
	var class string



	fmt.Printf("请输入学生id：")
	fmt.Scanf("%d\n",&sid)
	fmt.Printf("请输入学生姓名：")
	fmt.Scanf("%s\n",&name)
	fmt.Printf("请输入学生班级：")
	fmt.Scanf("%s\n",&class)
	fmt.Printf("请输入学生语文成绩：")
	fmt.Scanf("%d\n",&chinese)
	fmt.Printf("请输入学生数学成绩：")
	fmt.Scanf("%d\n",&math)
	fmt.Printf("请输入学生英语成绩：")
	fmt.Scanf("%d\n",&english)

	// make函数
	var score = make(map[string]int)   // 默认 开辟空间
	score["chinese"] = chinese
	score["math"] = math
	score["english"] = english
	score["average"] = (english+chinese+math)/3
	// 学生结构体对象
	stu:=newStudent(sid,name,class,score)

	return stu


}


func main() {

	manager:=&StuManager{}

	for true{
		//  展示菜单函数
		show_menu()

		// 引导用户输入选择
		var choice int
		fmt.Printf("请用户输入选择：")
		fmt.Scanf("%d\n",&choice)

		switch choice {
			case 1: manager.add_stu()
			case 2: manager.select_all_stus()
			case 3: manager.edit_stu()
			case 4: manager.score_sort()
			case 5: os.Exit(0)
		}
	}


}
