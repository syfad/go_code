package main

import "fmt"

//func main()  {

	//var info map[string]string
	//info = map[string]string{"name": "yuan", "age": "18","gender":"male"}
	//
	//fmt.Println(info)
	//fmt.Println(len(info))
	//fmt.Println(info["name"])  // 取某key值操作
	//fmt.Println(info["age"])


	// 访问方式:
	//info := map[string]string{"name": "yuan", "age": "18","gender":"male"}

	//for k,v :=range info{
	//	fmt.Println(k,v)
	//}

	//val,is_exist := info["names"]
	//fmt.Println(val,is_exist)

	// 修改
	//info["name"] = "alvin"
	//fmt.Println(info)
    //// 查询
	//delete(info,"age")
	//fmt.Println(info)

//}

func main()  {
	var info map[string]string
	info = map[string]string{"name": "yuan", "age": "18","gender":"male"}

	fmt.Println(info["name"])
	fmt.Println(info["age"])

}
