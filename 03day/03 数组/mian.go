package main

func main()  {


   // 声明  var 数组名 [长度] 元素的数据类型

	var arr [5] int
	//fmt.Println(arr) // [0,0,0,0,0]


	// 初始化方式1
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	// fmt.Println(arr)   // [1 2 3 4 5]
	// fmt.Println(arr[3])  // 4
	// fmt.Println(arr[6])  // error
	// fmt.Println(arr[-1])  // error


	// 初始化方式2
	//var arr2 = [3]string{"张三","李四","王五"}
	//fmt.Println(arr2)   // [张三 李四 王五]
	//fmt.Printf("%T",arr2)  // [3]string

	//var s1 [2] int  // [0,0,0]
	//var s2 [2] int  // [0,0]
	//fmt.Println(s1==s2)

	// 不定长元素
	//var arr3 = [...]string{"张三","李四","王五"}
	//var arr4 = [...]string{"李四","张三","王五"}
	//
	//fmt.Println(arr3)
	//fmt.Println(len(arr3))
	//fmt.Println(arr3 == arr4)


	// 索引设值
	//var arr5 = [5]int{0:100,2:22}
	//fmt.Println(arr5)


	// 数组占用的一整块连续的空间
	//fmt.Printf("%p\n",&arr[0])
	//fmt.Printf("%p\n",&arr[1])
	//fmt.Printf("%p\n",&arr[2])
	//fmt.Printf("%p\n",&arr[3])
	//fmt.Printf("%p\n",&arr[4])


	//  遍历数组
	// var arr6 = [...]string{"张三","李四","王五"}
	//for i:=0;i <len(arr6);i++ {
	//	fmt.Println(arr6[i])
	//}

	//for index,val :=range arr6{
	//	fmt.Println(index,val)
	//}


	//var arr7 [2][2]int
	//fmt.Println(arr7)   //[[0 0] [0 0]]
	//arr7[0][0] =1
	//arr7[0][1] =2
	//fmt.Println(arr7)
	//
	//var arr8 = [3][4]int{
	//	{0, 1, 2, 3} ,   //  第一行索引为 0
	//	{4, 5, 6, 7} ,   //  第二行索引为 1
	//	{8, 9, 10, 11} ,  //  第三行索引为 2
	//}
	////
	////fmt.Println(arr8) //[[0 1 2 3] [4 5 6 7] [8 9 10 11]]
	//
	//
	//for _,v :=range arr8{
    //    for i,v2 :=range v{
    //         fmt.Println(i,v2)
	//	}
	//}







}
