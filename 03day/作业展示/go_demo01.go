package main

import "fmt"

func BubbleSort(values []int) [] int {

	vLen := len(values)
	for i := 0; i < vLen-1; i++ {
		for j := 0; j < vLen-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}

	}

	return values
}


func main() {

	s:=[]int{10, 20, 32, 43, 123, 43, 22, 99, 9, 221,19}

	fmt.Println(BubbleSort(s))

}

/*



---i=0----
[20,10,35,11,14]

[10,20,35,11,14]
[10,20,35,11,14]
[10,20,11,35,14]
[10,20,11,14,35]


--i=1---
[10,20,11,14,35]

[10,20,11,14,35]
[10,11,20,14,35]
[10,11,14,20,35]


--i=2---
[10,11,14,20,35]

[10,11,14,20,35]
[10,11,14,20,35]

--i=3--
[10,11,14,20,35]



func bubble_sort(s []int){

       for i:=0;i<len(s)-1;i++{

            for j:=0;j<len(s)-i-1;j++{
                 if s[j]>s[j+1]{
                         s[j],s[j+1]=s[j+1],s[j]
                     }

            }

       }

}

len=5
i=0     4
i=1     3
i=2     2
i=3     1


bubble_sort([]int{12,34,4,56,7})





















 */