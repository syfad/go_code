package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//var x=10           // '10'
	//var name = "hi"    // '"hi"'
	//var b = true

	//json_str,_:=json.Marshal(x)
	//json_str2,_:=json.Marshal(name)
	//json_str3,_:=json.Marshal(b)
    //fmt.Printf("%#v\n",string(json_str))  // "10"
    //fmt.Printf("%#v\n",string(json_str2))  // "\"hi\""
    //fmt.Printf("%#v\n",string(json_str3))  // "true"

    // 结构体序列化

	type addr struct {
		Province string
		City string
	}
	type stu struct {
		Name string `json:"name"`   // tag
		Age int    // 隐私字段
		Addr addr
	}

	// var stu01 = stu{Name:"yuan", age:18, Addr:addr{Province:"Hebei",City:"beijing"}}
    // fmt.Printf("%#v",stu01)

	// json_str4,_:=json.Marshal(stu01)
    // fmt.Printf("%#v",string(json_str4))


	// 反序列化
	var data = `{"name":"yuan","Age":18,"Addr":{"Province":"Hebei","City":"beijing"}}`
	var stu02  stu

	json.Unmarshal([]byte(data),&stu02)
	fmt.Printf("%#v",stu02)

}
