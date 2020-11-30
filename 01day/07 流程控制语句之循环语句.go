package main

import "fmt"

func main() {




	//  条件循环  循环三要素: 初始值,条件,增量或减量
    //for  i := 1;i<=1000;i++{
	//	fmt.Println(i)
	//}
	//
	//for j := 1000;j>0;j--{
	//	fmt.Println(j)
	//}


	// 无限循环
	/*
	for true{
          fmt.Printf("OK")// 循环体
	}

	fmt.Printf("over!")

	 */


	// for循环方式2
	//count := 0
	//for count<=1000 {
	//	fmt.Println(count)
	//	count++
	//}
	//
	//for count<10{
	//	fmt.Println(count)
	//	count ++
	//}


	// 遍历循环 range
	// 遍历字符串
	//hello := "HelloYuan"
	//for i, v := range hello {
	//	fmt.Printf("%d,%c\n", k, v)
	//}


	//hello := "HelloYuan"
	//for i:=0;i<len(hello);i++{
	//	fmt.Println(i)
	//	fmt.Printf("%c",hello[i])
	//}


	// 嵌套循环
	//hello := "HelloYuan"
	//for i, v := range hello{
    //   fmt.Println(i,v)
    //   for j :=0;j<3;j++{
    //   	fmt.Printf("OK")
	//   }
	//
	//}


    // 关联嵌套
	//for i :=1;i<7;i++{
	//
	//	for j :=0;j<i;j++{
    //        fmt.Printf("*")
	//
	//	}
	//	fmt.Printf("\n")
	//
	//}


	// break

	//for i := 1; i < 100; i++ {
	//
	//	fmt.Println(":",i)
	//	if i == 14{
	//		// break   //  退出整个循环
	//		continue   //  退出当次循环
	//	}
	// }

/*	for i :=1;i<7;i++{

		for j :=0;j<i;j++{

	       if j==3{
	       	break
		   }
		   fmt.Printf("*")

		}
		fmt.Printf("\n")

	}

	fmt.Printf("over!")


*/

	/*

	1 1 2 3 5 8 13 21


	 */
    k1 :=1
    k2:=1
    z:=0
    for i:=1;i<=60;i++{
    	z=k1+k2
    	//z=sum
    	k1=k2
    	k2=z

	}
	fmt.Println(k1,k2,z)
}

