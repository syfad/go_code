package main

func main() {


	//  值类型  :  int float string bool 数组 struct



	//var a int   //int类型默认值为 0
	//var b string    //string类型默认值为""
	//var c bool      //bool类型默认值为false
	//var d [2]int    //数组默认值为[0 0]
	//fmt.Println(&a) //默认已经分配内存地址，可以使用&来取内存地址


	//var a =10
	//b := a    // 将a对应的值10拷贝了一份赋值给了b
	//a = 101
	//fmt.Printf("a的值是%v，a的内存地址是%p\n",a,&a)   // 10  oxxxxxxxx
	//fmt.Printf("b的值是%v，b的内存地址是%p\n",b,&b)   // 101 oxxxxxxxx
	//
	//
	////数组的赋值
	//var c =[3]int{1,2,3}    //定义一个长度为3的int类型的数组
	//d := c      //将数组c拷贝赋值给d
	//d[1] = 100  //修改数组d中索引为1的值为100
	//fmt.Printf("c的值是%v，c的内存地址是%p\n",c,&c)   //c的值是[1 2 3]，c的内存地址是0xc42000a180
	//fmt.Printf("d的值是%v，d的内存地址是%p\n",d,&d)   //d的值是[1 100 3]，d的内存地址是0xc42000a1a0


	// 引用类型: 切片  map


	//var a = []int{1,2,3,4,5}
	//b := a      //此时a，b都指向了内存中的[1 2 3 4 5]的地址
	//b[1] = 10   //相当于修改同一个内存地址，所以a的值也会改变
	//fmt.Printf("%p",&a)
	//fmt.Printf("%p",&b)
	//
	//
	//
	//c := make([]int,5,5)    //切片的初始化
	//fmt.Println(c)  // [0 0 0 0 0]
	//
	//copy(c,a)   //将切片a copy到 c
	//fmt.Println(c)
	//
	//c[1] = 20   //copy是值类型，所以a不会改变
	//fmt.Printf("a的值是%v，a的内存地址是%p\n",a,&a)  // [1,10,3,4,5]   oxxxxxx
	//fmt.Printf("b的值是%v，b的内存地址是%p\n",b,&b)  // [1,10,3,4,5]   oxxxxxx
	//fmt.Printf("c的值是%v，c的内存地址是%p\n",c,&c)  // [1,20,3,4,5]   ox
	//
	//d := &a     //将a的内存地址赋值给d，取值用*d
	//a[1] = 11
	//fmt.Printf("d的值是%v，d的内存地址是%p\n",*d,d)  // [1,11,3,4,5]
	//fmt.Printf("a的值是%v，a的内存地址是%p\n",a,&a)  // [1,11,3,4,5]


	// make函数

	//var arr [] int  // var arr [2] int
	//arr[0] = 1   // arr:= []int{1,2,3}
	//fmt.Println(arr)

	//var m map[string]string  // var m = map[string]string{}
	//m["k1"] = "v1"
	//fmt.Println(m)

	//a := make([]int, 2)
	//a[0] =1
    //fmt.Println(a)  // [1 0]


    // append追加功能 ,构建切片首位置添加
	//var a = []int{1,2,3}  // [1,2,3]
	//temp:=[]int{0}  // [添加到首位置的元素]
	//a = append(temp, a...)   // [0,1,2,3]
	//
	//a = append([]int{-3,-2,-1}, a...)


	// append构建一个向某一个位置插入元素或者切片的方法
	//var a =[]int{1,2,3,4,5}
	//
	//// 向索引为2的位置插入100
	//temp01:=a[:2]   // [1,2]
	//temp02:=a[2:]   // [3,4,5]
	//
	//temp03:=[]int{100}
	//temp04:=append(temp03,temp02...)   // [100,3,4,5]
	//
	//ret:=append(temp01,temp04...)
	//fmt.Println(ret)   //  [1 2 100 3 4 5]

	// a = append(a[:i], append([]int{x}, a[i:]...)...) //   [1,2,100,3,4,5]


	// 扩容机制
	// 情况1:容量不足
	//slice := []int{10, 20, 30, 40}
	//newSlice := append(slice, 50)
	//newSlice[1] += 10
	//
	//fmt.Println(slice) // [10,20,30,40]
	//fmt.Println(newSlice) // [10,30,30,40,50]

	// 情况2: 容量够
	//array := []int{10, 20, 30, 40}
	//slice := array[0:2] // [10,20]          array[0:2:2]
	//newSlice := append(slice, 50)
	//
	//newSlice[1] += 10
	//
	//fmt.Println(slice) // [10,30]
	//fmt.Println(newSlice) // [10,30,50]
	//fmt.Println(array)  // [10,30,50,40]
	//
    //fmt.Printf("%p\n",array)
    //fmt.Printf("%p\n",slice)
    //fmt.Printf("%p\n",newSlice)


	// 面试题:

	//arr := [4]int{10, 20, 30, 40}
	//slice := arr[0:2]  // slice =[10,20]  cap=4
	//testSlice1 := slice  // 引用类型
	//testSlice2 := append(append(append(slice, 1),2),3)
	//slice[0] = 11
	//fmt.Println(slice)  // [11,20]
	//fmt.Println(testSlice1)  // [11,20]
	//fmt.Println(testSlice2)  // [10,20,1,2,3]
	//fmt.Println(arr)









}
