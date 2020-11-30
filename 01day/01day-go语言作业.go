package main

import "fmt"

//第一题
//输入周几,返回对应提示
func main(){
	for true {
		var weekday int
		fmt.Printf("请输入今天是星期几？\n")
		fmt.Scanf("%d", &weekday)

		if (weekday == 6) || (weekday == 7) {
			fmt.Printf("今天是周%d, 周末愉快\n", weekday)
		} else if weekday >= 1 && weekday <= 6 {
			fmt.Printf("工作日愉快\n")
		} else {
			fmt.Printf("请输入位于1 ~ 7之间的数字\n")
		}
	}
}

//第二题
//要求用户输入一个年份和一个月份，判断（要求使用嵌套的if…else和switch分别判断一次）该年该月有多少天。

func main()  {
	var (
		year int
		month int
		//days int
	)
	fmt.Printf("please input a year,month:\n")
	fmt.Scanf("%d,%d", &year,&month)

	if (month != 2){
		if month == 4 || month == 6 || month == 9 || month == 11{
			fmt.Printf("该月有30天")
		}else {
			fmt.Printf("该月有31天")
		}
	}else {
		if (((year % 4) == 0 && (year % 100) != 0) || (year % 400) == 0 ){
			fmt.Printf("该月有29天")
		}else {
			fmt.Printf("该月有28天")
		}
	}

}


//第三题
//一只猴子第一天摘下若干个桃子，当即吃了一半，不过瘾，又多吃了一个
//第二天早上又将剩下的桃子吃掉一半，又多吃了一个。以后每天早上都吃了前
//一天剩下的一半零一个，第10天早上想再吃时，发现只剩下一个桃子了。请问
//猴子第一天一共摘了多少个桃子？
func main()  {
	var s int
	s = 1
	for i := 1;i<9;i++{
		s=(s+1)*2
	}
	fmt.Printf("第一天摘了%d", s)
}