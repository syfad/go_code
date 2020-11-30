package main
//
import (
	"fmt"
	"os"
)

type student struct {
	sid   int
	name  string
	class string
	score map[string]int
}

func newStudent(sid int, name string,class string,score map[string]int)(*student){
  return &student{
	   sid:   sid,
	   name:  name,
	   class: class,
	   score: score,
  }
}

func add_stu_info()  {
	new_stu:=get_input_stu()
	all_students = append(all_students, new_stu)
}

func edit_stu_info()  {
	edit_stu:=get_input_stu()
	for i,stu :=range all_students{
		if stu.sid == edit_stu.sid{
			all_students[i] = edit_stu
			break
		}

	}
}

func select_allstus_info()  {
   for _,stu :=range all_students{
       fmt.Printf("学号：%5d，姓名：%-5s，班级，%s，语文成绩：%5d，数学成绩%5d，英文成绩%-5d，平均成绩%5d\n",stu.sid,stu.name,stu.class,stu.score["chinese"],stu.score["math"],stu.score["english"],stu.score["average"])
	}
}

func score_sort()  {

	vLen := len(all_students)
	for i := 0; i < vLen-1; i++ {
		for j := 0; j < vLen-i-1; j++ {
			if all_students[j].score["average"] > all_students[j+1].score["average"] {
				all_students[j],all_students[j+1] = all_students[j+1],all_students[j]
			}
		}

	}
	select_allstus_info()
}


func show_menu(){
	fmt.Println("欢迎来到学生成绩管理系统！")
	fmt.Println("1、添加学生成绩")
	fmt.Println("2、查询所有学生成绩")
	fmt.Println("3、编辑学生成绩")
	fmt.Println("4、成绩从高到低排序！")
	fmt.Println("4、退出程序！")

}

func get_input_stu()*student{
	var sid,chinese,math,english int
	var name string
	var class string
	var scores map[string]int
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
	scores = make(map[string]int)
	scores["chinese"]=chinese
	scores["math"]=math
	scores["english"]=english
	scores["average"]=(chinese+math+english)/3
	return newStudent(sid,name,class,scores)
}


var all_students []*student


func main() {

	for true {

		show_menu()
		var choice int
		fmt.Printf("请用户输入选择：")
		fmt.Scanf("%d\n",&choice)

		switch choice {
		case 1:add_stu_info()
		case 2:select_allstus_info()
		case 3:edit_stu_info()
		case 4:score_sort()
		case 5:os.Exit(0)

		}
	}


}




/*

[20,10,35,11,14]


*/