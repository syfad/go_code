package main

import (
	"encoding/json"
	"fmt"
)

type addr struct {
	Province,City string
}

type stu struct {
	Name string
	Age int
	Addr addr
}

func main() {
	var stu01 = stu{Name: "syf",Age: 18, Addr: addr{Province: "北京", City: "朝阳"}}

	json_data, err :=json.Marshal(stu01)
	if err  != nil{
		fmt.Println("json解析失败，原因", err)
		return
	}
	data := string(json_data)
	fmt.Println(data)
	var stu02 stu

	json.Unmarshal([]byte(data), &stu02)
	fmt.Println(stu02.Addr,stu02.Name)

}