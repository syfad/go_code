package main

import "fmt"

//Animal 动物
type Animal struct {
	name string
	age  int
}

func (a *Animal) eat() {
	fmt.Printf("%s eating！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int
	Animal //通过嵌套匿名结构体实现继承
}

type Cat struct {
	color string
	pinzhong string
	Animal
}

func(c Cat) catch_mouse()  {
	fmt.Printf("%v正在抓老鼠!",c.name)
}

func(c Cat) eat()  {
	fmt.Printf("%v正在吃老鼠!",c.name)
}

func (d *Dog) bark () {
	fmt.Printf("%s barking ~\n", d.name)
}

func main() {

	a1:=Animal{ //注意嵌套的是结构体指针
		name: "大黄",
		age:12,
	}
	d1 := Dog{
		Feet: 4,
		Animal: a1,
	}

	fmt.Println(d1.Feet)
	fmt.Println(d1.name)   // 查找顺序:结构体对象查找字段名称先去自身结构体中查找,如果没有,去匿名字段查找
	d1.bark()

	//c1:=Cat{
	//	color: "red",
	//	pinzhong: "加菲",
	//	Animal:Animal{
	//		name:"花花",
	//		age:3,
	//	},
	//}
	//
	////fmt.Println(c1.color)
	//c1.catch_mouse()
	//c1.eat()

}