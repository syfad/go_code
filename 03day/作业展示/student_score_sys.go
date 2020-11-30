package main

import (
	"fmt"
	"os"
)

var stus_scores = make(map[string] map[string] int)


func select_stu_score(){

	var select_name string
	fmt.Printf("请输入查询学生的姓名:")
	fmt.Scanf("%s",&select_name)
	yuwen_score := stus_scores[select_name]["yuwen_score"]
	shuxue_score := stus_scores[select_name]["shuxue_score"]
	yingyu_score := stus_scores[select_name]["yingyu_score"]
	fmt.Printf("姓名：%s: 语文成绩%d;数学成绩%d;英文成绩%d",select_name,yuwen_score,shuxue_score,yingyu_score)

}

func selectall_stus_score(){

	fmt.Printf(`

 学生成绩一览表：
 -----------------------------------------------------------------------
	姓名		语文	数学	英语	平均分
 -----------------------------------------------------------------------
`)

	for name,scores :=range stus_scores{
		aver_acore:=(scores["yuwen_score"]+scores["shuxue_score"]+scores["yingyu_score"])/3
		fmt.Printf("\t%-4s\t\t%-4d\t%-4d\t%-4d\t%-4d\n",name,scores["yuwen_score"],scores["shuxue_score"],scores["yingyu_score"],aver_acore)
	}
	fmt.Printf(`
 -----------------------------------------------------------------------
`)
}


func add_stu_score(){


	var name string
	var yuwen_score,shuxue_score,yingyu_score int
	var stu_score = make(map[string] int)

	fmt.Printf("请输入学生的姓名:\n")
	fmt.Scanf("%s",&name)
	fmt.Printf("请输入学生的语文成绩:\n")
	fmt.Scanf("%d",&yuwen_score)
	fmt.Printf("请输入学生的数学成绩:\n")
	fmt.Scanf("%d",&shuxue_score)
	fmt.Printf("请输入学生的英语成绩:\n")
	fmt.Scanf("%d",&yingyu_score)


	stu_score["yuwen_score"] = yuwen_score
	stu_score["shuxue_score"] = shuxue_score
	stu_score["yingyu_score"] = yingyu_score

	stus_scores[name]=stu_score

}


func update_stu_score(){

	var update_name string
	fmt.Printf("请输入修改的学生的姓名:")
	fmt.Scanf("%s",&update_name)


}


func delete_stu(){
	var delete_name string
	fmt.Printf("请输入删除的学生的姓名:")
	fmt.Scanf("%s",&delete_name)
  	delete(stus_scores,delete_name)
	fmt.Printf("删除成功！")
}


func paixu(){
   // 二分查找排序

}

func main() {

	var choice int

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
		fmt.Scanf("%d",&choice)

		switch choice {

		    case 1: add_stu_score()
			case 2: select_stu_score()
			case 3: selectall_stus_score()
			case 4: update_stu_score()
			case 5: delete_stu()
			case 6: paixu()
			case -1:os.Exit(1)

			default:fmt.Printf("请输入正确操作数字")

		}

	}


}