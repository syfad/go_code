package main

import "fmt"

type Maoke interface {
	bark
	eat
}

type bark interface {
	jiaohuan()
}

type eat interface {
	eat()
}

type Dog struct {
	name string
}

func (d *Dog)eat(){
	fmt.Printf("%s正在吃屎...",d.name)
}

func (d *Dog)jiaohuan(){
	fmt.Printf("%s正在叫唤...",d.name)
}

type xxx interface {

}

func foo(a interface{})  {
     fmt.Printf("%#v\n",a)
}

func main() {


	// 值类型和指针类型
    //var b bark
    // alex := &Dog{name: "alex"}
    //alex := &Dog{name: "alex"}   // 方法如果是一个值类型,那么结构体对象和结构体的指针变量都可以调用该方法
    //b= alex

    // b.jiaohuan()

    //x:=10
	//
	//// 接口和类型的关系 :   一个接口可以对应多个类型;一个类型也可以对应多个接口
	//alex.eat()
	//
    //// 接口嵌套
    //var m Maoke
    //m=alex
    //fmt.Println(m) // &{alex}


    // 口接口: 接口为空:

	//var x interface{}
    //// x可以接收一切类型
    //x=10
    //x="yuan"
    //x=[]int{12,324,5456}
	//
    //m:=make(map[string]interface{})
    //m["name"]="yuan"
    //m["age"]=23
    //fmt.Println(m)
	//
	//foo(123)
	//foo("hi")
	//foo(m)

    //

    var y interface{}
    y="123"
    var v,is_ok=y.(string)
    fmt.Println(v,is_ok)
	switch y.(type) {
	case string:fmt.Printf("string")
	case int:fmt.Printf("int")
	case bool:fmt.Printf("bool")

	}





}
