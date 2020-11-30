package main

import (
	"bufio"
	"fmt"
	"os"
)

//结构体
type student struct {
	sid   int
	name  string
	class string
	score map[string]int
	pwd	string
}

//模拟构造函数
func newStudent(sid int, name string, class string, score map[string]int, pwd string) *student {
	return &student{
		sid:   sid,
		name:  name,
		class: class,
		score: score,
		pwd: pwd,
	}
}

//设计学生管理结构
type stu_manage struct {
	stu_list []*student
}

//var all_students []*student
func (s *stu_manage) add_stu_info() {
	fmt.Printf("添加学生成绩")
	new_stu := get_input_stu()
	s.stu_list = append(s.stu_list, new_stu)

	file, err := os.OpenFile("./login.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil{
		fmt.Println("open file failed, err", err)
		return
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(new_stu.namenew_stu.pwd)
	write.WriteString(new_stu.pwd)
	write.Flush()

}

func (s *stu_manage) select_all_stu() {
	fmt.Printf("查询所有学生成绩！")
	fmt.Println(s.stu_list )
	for _,stu :=range s.stu_list{
		fmt.Println(stu.sid,stu.class,stu.score,stu.name)
	}
}

func (s *stu_manage) edit_stu() {
	new_stu := get_input_stu()

	for i := 0; i < len(s.stu_list); i++ {
		fmt.Println(s.stu_list[i])
		if s.stu_list[i].sid == new_stu.sid {
			s.stu_list[i] = new_stu
		}
	}
	fmt.Println("编辑学生成绩")

}



func get_input_stu() *student {
	var sid, chinese, math, english int
	var name string
	var class string
	var scores map[string]int
	fmt.Printf("请输入学生id：")
	fmt.Scanf("%d\n", &sid)
	fmt.Printf("请输入学生姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Printf("请输入学生班级：")
	fmt.Scanf("%s\n", &class)
	fmt.Printf("请输入学生语文成绩：")
	fmt.Scanf("%d\n", &chinese)
	fmt.Printf("请输入学生数学成绩：")
	fmt.Scanf("%d\n", &math)
	fmt.Printf("请输入学生英语成绩：")
	fmt.Scanf("%d\n", &english)
	scores = make(map[string]int)
	scores["chinese"] = chinese
	scores["math"] = math
	scores["english"] = english
	scores["average"] = (chinese + math + english) / 3
	return newStudent(sid, name, class, scores)
}

func show_menu() {

	fmt.Println("欢迎来到学生成绩管理系统！")
	fmt.Println("1、添加学生成绩")
	fmt.Println("2、查询所有学生成绩")
	fmt.Println("3、编辑学生成绩")
	fmt.Println("4、成绩从高到低排序！")
	fmt.Println("5、退出程序！")

}




func main() {
	manage := &stu_manage{}


	for true {
		show_menu()
		var choose_id int
		fmt.Scanf("%d", &choose_id)

		switch choose_id {
		case 1:
			manage.add_stu_info()
			manage.write_stu()
		case 2:
			manage.select_all_stu()
		case 3:
			manage.edit_stu()

		}
	}


}
