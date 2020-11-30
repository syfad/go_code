package main

//type Student2 struct {
//	string
//	age  int
//}



//  Student结构体下方法
//func (s Student) print_stu_info(){
//   fmt.Printf("学生姓名:%s,年龄:%d",s.name,s.age)
//}
//
//func (s Student)setStuAge(newAge int){
//     s.age=newAge
//     fmt.Println("方法中",s.age)
//}

//func (s *Student)setStuAge2(newAge int){
//	s.age=newAge
//	fmt.Println("方法中",s.age)
//}

//MyInt 将int定义为自定义MyInt类型
//type MyInt int
//
////SayHello 为MyInt添加一个SayHello的方法
//func (m MyInt) SayHello() {
//	fmt.Println("Hello, 我是一个int。")
//}

// 匿名字段

type Book struct {
	string
	pubulish string
}

func main() {

	//var i  = 10
	//var x =&i  // x :指针类型 *int
	//fmt.Println(*x) // 10 将x这个指针变量存储的地址对应值取出来



	// 接收器
	//var stu01 = Student{name:"yuan",age:18}
	//stu01.print_stu_info()
	////var stu02 = Student{name:"alvin",age:22}
	////stu02.print_stu_info()
	//
    ////  值类型和指针类型
	//
    //stu01.setStuAge2(1000)
	//fmt.Println("mian中:",stu01.age)


	// 为int类型添加方法

	//var x  MyInt
	//x=10
	//x.SayHello()


	// 匿名字段

	//b:=&Book{"西游记","北京出版社"}
	//
    //fmt.Println(b.string)
    //fmt.Println(b.pubulish)


    /*

    class Animal():
         def __init__(self,name,age):  // 构造函数
              self.name = name
              self.age = age


    // 实例初始化
    alex = Animal("alex",1000)

   // go
    type Animal struct {
		name string
		age  int
	}

     // go语言的实例初始化
     var alex = &Animal{"name":"alex","age":1000}


     // 模拟构造函数

     func NewAnimal(name string,age  int)(*Animal){

            return &Animal{
                  "name":name,
                  "age":age}
     }

     var alex= NewAnimal("alex",1000)



    //

    class  A:
        属性=属性值
        属性=属性值
        属性=属性值
		def  foo1():pass
		def  foo2():pass
		def  foo3():pass
		def  foo4():pass


    class  B:
        属性=属性值
		属性=属性值
		属性=属性值
		def  bar1():pass
		def  bar2():pass
		def  bar3():pass
		def  bar4():pass

-----------------------------------

def  foo1():pass
def  foo2():pass
def  foo3():pass
def  foo4():pass

def  bar1():pass
def  bar2():pass
def  bar3():pass
def  bar4():pass

属性=属性值
属性=属性值
属性=属性值

属性=属性值
属性=属性值
属性=属性值

// go语言实现:接收器



     */
}
