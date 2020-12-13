package main

import "fmt"

func main() {
	s := "上海自来水来自海上"
	is_huiwen1(s)
}

//自己想的方法
func is_huiwen(str1 string) {
	var r = []rune(str1)

	for i := 0; i < len(r); i++ {
		x := 0
		j := (len(r) - 1)
		if r[x] == r[j] {
			fmt.Println("is huiwen")
		} else {
			fmt.Println("no huiwen")
		}
	}
}

//标准答案
func is_huiwen1(str1 string) bool {
	if len(str1) == 0 || len(str1) == 1 {
		fmt.Println("长度不够")
		return false
	}
	var r = []rune(str1)
	i, j := 0, len(r)-1
	for i < j {
		if r[i] == r[j] {
			i++
			j--
			//fmt.Println(i,j)
		} else {
			fmt.Println("no huiwen")
			return false
		}
	}
	fmt.Println("is huiwen")
	return true

}
