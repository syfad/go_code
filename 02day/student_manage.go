package main

import (
	"fmt"
	"os"
)

var stu_manage = make(map[string]map[string]int)

//stu_manage := ["syf":["yuwen":12, "shuxue":56, "yingyu": 66]]

func add_stu() {
	var stu_name string
	var yuwen, shuxue, yingyu int
	var stu_score = make(map[string]int)

	fmt.Printf("请输入学生姓名：")
	fmt.Scanf("%s", &stu_name)
	fmt.Printf("请输入语文成绩：")
	fmt.Scanf("%d", &yuwen)
	fmt.Printf("请输入数学成绩：")
	fmt.Scanf("%d", &shuxue)
	fmt.Printf("请输入英语成绩：")
	fmt.Scanf("%d", &yingyu)

	stu_score["yuwen"] = yuwen
	stu_score["shuxue"] = shuxue
	stu_score["yingyu"] = yingyu
	stu_manage[stu_name] = stu_score
	//fmt.Println(stu_manage)

}

func select_stu_manage() {
	var stu_name string
	fmt.Printf("请输入学生姓名：")
	fmt.Scanf("%s", &stu_name)
	fmt.Printf("学生姓名：%s,语文成绩：%d, 数学成绩：%d, 英语成绩：%d", stu_name, stu_manage[stu_name]["yuwen"], stu_manage[stu_name]["shuxue"], stu_manage[stu_name]["yingyu"])
	fmt.Printf("\n=============================================================\n\n")
}
func select_all_stu_manage() {
	//var stu_name string
	//fmt.Printf("请输入学生姓名：")
	//fmt.Scanf("%s", &stu_name)

	fmt.Printf(`

 学生成绩一览表：
 -----------------------------------------------------------------------
姓名		语文	数学	英语	平均分
 -----------------------------------------------------------------------
`)
	for k, v := range stu_manage {
		fmt.Println(k, "  ", v["yuwen"], v["shuxue"], v["yingyu"], (v["yuwen"]+v["shuxue"]+v["yingyu"])/3)
	}

}

func delete_stu_manage() {
	var stu_name string
	fmt.Printf("请输入学生姓名：")
	fmt.Scanf("%s", &stu_name)
	if len(stu_manage[stu_name]) != 0 {
		delete(stu_manage, stu_name)
		fmt.Printf("Delete Done")
	} else {
		fmt.Printf("学生姓名不存在，请重新选择！\n")
		fmt.Printf("======================================")
	}
}

func update_stu_manage() {
	var stu_name, subject string
	var score int
	fmt.Printf("请输入学生姓名:")
	fmt.Scanf("%s", &stu_name)
	fmt.Printf("学生姓名：%s,语文成绩：%d, 数学成绩：%d, 英语成绩：%d,\n", stu_name, stu_manage[stu_name]["yuwen"], stu_manage[stu_name]["shuxue"], stu_manage[stu_name]["yingyu"])
	fmt.Printf("请输入需要修改的科目名（yuwen，shuxue，yingyu）")
	fmt.Scanf("%s", &subject)
	fmt.Printf("请输入需要修改的本科目的分数:")
	fmt.Scanf("%d", &score)
	stu_manage[stu_name][subject] = score
	//fmt.Println(stu_manage[stu_name])
	//fmt.Println(stu_manage[stu_name][subject])
}

func main() {
	var num int
	for true {
		fmt.Printf(`
	***************************学生成绩管理系统*****************
	  	1.添加学生成绩信息
		2.查询学生成绩信息
		3.查询所有学生成绩信息
		4.修改学生成绩信息
		5.删除学生成绩信息
		6.成绩排序
	  	-1:退出系统
	************************************************************
		请输入您的选择:

	`)
		fmt.Scanf("%d", &num)
		switch num {
		case 1:
			add_stu()
		case 2:
			select_stu_manage()
		case 3:
			select_all_stu_manage()
		case 4:
			update_stu_manage()
		case 5:
			delete_stu_manage()
		case 6:
			fmt.Printf("正在开发中，敬请期待！")
		case -1:
			os.Exit(0)
		default:
			fmt.Printf("请按指示输入正确的数字：")
		}
	}
}
