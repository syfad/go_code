package main

import "fmt"

//冒泡排序
func bubble_sort(s []int) []int{
	var lens = len(s)
	for i := 0; i < lens-1; i++ {
		for j := 0; j < lens-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}

func main() {
	sort := []int{10, 20, 32, 43, 123, 43, 22, 99, 9, 221,19}

	fmt.Println(bubble_sort(sort))

}
