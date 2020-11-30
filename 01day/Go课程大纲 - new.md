# Go周末课程大纲

## 一、Go基本介绍

<img src = '.\assets\QQ截图20201016174250.png' width= 500 height=300>

### 1.1、认识Go语言

*  go语言（或 Golang）是Google开发的开源编程语言，诞生于2006年1月2日下午15点4分5秒，于2009年11 月开源，2012年发布go稳定版

* go是非常年轻的一门语言，它的主要目标是“兼具Python 等动态语言的开发速度和C/C++等编译型语言的性能与安全性”

> 注意，编译语言和解释器语言的区别！

### 1.1、环境安装

**1.1.1、 下载SDK**

* 官网：https://golang.google.cn/

* go中文网：https://studygolang.com/dl

**1.1.2、安装和配置SDK**

![](J:\Go周末班\Go课程大纲\assets\图片1.png)

![](J:\Go周末班\Go课程大纲\assets\图片2.png)

![](J:\Go周末班\Go课程大纲\assets\图片3.png)

![](J:\Go周末班\Go课程大纲\assets\图片4.png)

```
执行命令 go version 查看安装go的版本
```

*  配置环境变量GOPATH和GOBIN

* 下面是sdk的安装位置，若没有，自己配一下

![](J:\Go周末班\Go课程大纲\assets\图片5.png)

* GOPATH可以自定义

![](J:\Go周末班\Go课程大纲\assets\图片6.png)



* GOPATH里面，手动创建如下3个目录

```
  src是放代码，放自己写的代码

  pkg：编译时生成的中间文件

  bin：GOBIN就是配到这个bin里面
```

![](J:\Go周末班\Go课程大纲\assets\图片7.png)

**1.1.3、安装开发工具（GoLand）**

下载地址：https://www.jetbrains.com/go/download/#section=windows

安装步骤：

![](J:\Go周末班\Go课程大纲\assets\图片8.png)



![](J:\Go周末班\Go课程大纲\assets\图片9.png)

![](J:\Go周末班\Go课程大纲\assets\图片10.png)

### 1.2、第一个Go程序

第一个Go程序的执行

```go

package main

import (
	"fmt"
)

func main() {

	fmt.Print("hello world!")

}
```

1、Package main is special. It defines a standalone executable program, not a library. Within package main the function main is also special—it’s where execution of the program begins. Whatever main does is what the program does.

`main package`不同于其它`library package`，它定义了一个可执行程序。其中的`main`函数即是可执行文件的入口函数。

2、Println与Printf的区别

```go
func main() {

	a := 10
	b := 20

	fmt.Println(a) //可以打印出字符串，和变量,有默认换行
	fmt.Println(b)
	fmt.Println("abc") //right
	fmt.Printf("%d",a) //默认不换行
	fmt.Println("%d",a) //error
	// fmt.Printf(a) //error

}
```

3、fmt包是Go语言的最重要的标准库之一，fmt包实现了类似C语言printf和scanf的格式化I/O。主要分为向外输出内容和获取输入内容两大部分。

```go
package main

import "fmt"

func main() { 
    //  { 不能在单独的行上，很有趣的设定
	var (
		name    string
		age     int
	)
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("扫描结果 name:%s age:%d\n", name, age)
}
```

4、执行：

```
（1）方式1：编译go build

使用`go build`

1. 在项目目录下执行`go build`
2. 在其他路径下执行`go build`， 需要在后面加上项目的路径（项目路径从GOPATH/src后开始写起，编译之后的可执行文件就保存在当前目录下）
3. `go build -o hello.exe`‘’
4. 执行hello.exe即可

（2）方式2：go run 文件名

```

## **二、** **基础语法**

### 2.1 命名规则

go语言中的函数名、常量名、类型名、语句标号和包名等所有的命名，都遵循一个简单的命名规则，必须以一个字母或下划线_开头，后面可以跟任意数量的字母、数字或下划线，且go语言中大小写字母认为是不同的。

go语言中有25个关键字，不能用于自定义名字

```go
break        default          func           interface         select
case         defer            go             map               struct
chan         else             goto           package           switch
const        fallthrough      if             range             type
continue     for              import         return            var
```

还有30多个预定义的名字，用于内建的常量、类型和函数

```go
内建常量:
    true false iota nil
内建类型:
    int int8 int16 int32 int64
    uint uint8 uint16 uint32 uint64 uintptr
    float32 float64 complex128 complex64
    bool byte rune string error
内建函数:
    make len cap new append copy close delete
    complex real imag
    panic recover
```

### 2.2 变量

#### 2.2.1、声明变量的语法格式：

```go
var 变量名 变量类型 // 变量的声明必须使用空格隔开
```

```go
func main() {

	// 声明一个变量
	var v1 int
	var v2 int

	// 声明多个变量
	var v3,v4 int

	// 声明多个变量another way
	var (
       v5 int
       v6 bool
	   v7 string
	)
	
	fmt.Println(v1,v2,v3,v4,v5,v6,v7)  //任何类型变量声明不使用报错
}
```

#### 2.2.2 变量赋值

```go
func main() {

	// 方式1:先声明再赋值
	var name string
	name = "yuan"
	// 方式2: 声明并赋值
	var age = 18

	// 声明并赋值简写
	gender := "male"

	// 多重赋值
	var v1,v2 int
	v1,v2 = 1,2

	// 匿名变量
	var v3,_ = 5,6  // _ 是占位的作用

	// 不能重复声明变量
	// var v3 = 100

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(gender)
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)

}
```

### 2.3、常量

相对于变量，常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值。 常量的声明和变量声明非常类似，只是把`var`换成了`const`，常量在定义的时候必须赋值。

```go
const pi = 3.1415
const e = 2.7182
```

声明了`pi`和`e`这两个常量之后，在整个程序运行期间它们的值都不能再发生变化了。

多个常量也可以一起声明：

```go
const (
    pi = 3.1415
    e = 2.7182
)
```

const同时声明多个常量时，如果省略了值则表示和上面一行的值相同。 例如：

```go
const (
    n1 = 100
    n2
    n3
)
```

上面示例中，常量`n1`、`n2`、`n3`的值都是100。

### 2.4、iota

`iota`是go语言的常量计数器，只能在常量的表达式中使用。

`iota`在const关键字出现时将被重置为0。const中每新增一行常量声明将使`iota`计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。

举个例子：

```go
const (
		n1 = iota //0
		n2        //1
		n3        //2
		n4        //3
	)
```

**几个常见的`iota`示例:**

使用`_`跳过某些值

```go
const (
		n1 = iota //0
		n2        //1
		_
		n4        //3
	)
```

`iota`声明中间插队

```go
const (
		n1 = iota //0
		n2 = 100  //100
		n3 = iota //2
		n4        //3
	)
	const n5 = iota //0
```

定义数量级 （这里的`<<`表示左移操作，`1<<10`表示将1的二进制表示向左移10位，也就是由`1`变成了`10000000000`，也就是十进制的1024。同理`2<<2`表示将2的二进制表示向左移2位，也就是由`10`变成了`1000`，也就是十进制的8。）

```go
const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
```

多个`iota`定义在一行

```go
const (
		a, b = iota + 1, iota + 2 //1,2
		c, d                      //2,3
		e, f                      //3,4
	)
```

### 2.4、注释

注释不会被编译，每一个包应该有相关注释。单行注释是最常见的注释形式，你可以在任何地方使用以 // 开头的单行注释。多行注释也叫块注释，均已以 /* 开头，并以 */ 结尾。如：

```go
// 单行注释
/*
   多行注释
*/
```

### 2.5、行分隔符

在 Go 程序中，一行代表一个语句结束。每个语句不需要像 C 家族中的其它语言一样以分号 ; 结尾，因为这些工作都将由 Go 编译器自动完成。如果你打算将多个语句写在同一行，它们则必须使用 ; 人为区分。这一点是借鉴的python语法特性。

```go
func main() {

	var name = "yuan";var age = 18 // 不推荐
	fmt.Println(name)
fmt.Println(age)  // 不报错但是不推荐

}
```

## 三、基本数据类型

| 序号 | 类型和描述                                                   |
| :--- | :----------------------------------------------------------- |
| 1    | **布尔型** 布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。 |
| 2    | **数字类型** 整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。 |
| 3    | **字符串类型:** 字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。 |

### 3.1、整型

整型分为以下两个大类： 按长度分为：int8、int16、int32、int64 对应的无符号整型：uint8、uint16、uint32、uint64

其中，`uint8`就是我们熟知的`byte`型，`int16`对应C语言中的`short`型，`int64`对应C语言中的`long`型。

|  类型  |                             描述                             |
| :----: | :----------------------------------------------------------: |
| uint8  |                  无符号 8位整型 (0 到 255)                   |
| uint16 |                 无符号 16位整型 (0 到 65535)                 |
| uint32 |              无符号 32位整型 (0 到 4294967295)               |
| uint64 |         无符号 64位整型 (0 到 18446744073709551615)          |
|  int8  |                 有符号 8位整型 (-128 到 127)                 |
| int16  |              有符号 16位整型 (-32768 到 32767)               |
| int32  |         有符号 32位整型 (-2147483648 到 2147483647)          |
| int64  | 有符号 64位整型 (-9223372036854775808 到 9223372036854775807) |

**特殊整型**

| 类型 |                          描述                          |
| :--: | :----------------------------------------------------: |
| uint | 32位操作系统上就是`uint32`，64位操作系统上就是`uint64` |
| int  |  32位操作系统上就是`int32`，64位操作系统上就是`int64`  |

我们可以借助fmt函数来将一个整数以不同进制形式展示。

```go
package main
 
import "fmt"
 
func main(){
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a)  // 10
	fmt.Printf("%b \n", a)  // 1010  占位符%b表示二进制
 
	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b)  // 77
 
	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c)  // ff
	fmt.Printf("%X \n", c)  // FF
}
```

### 3.2、浮点型

Go语言支持两种浮点型数：`float32`和`float64`。这两种浮点型数据格式遵循`IEEE 754`标准： `float32` 的浮点数的最大范围约为 `3.4e38`，可以使用常量定义：`math.MaxFloat32`。 `float64` 的浮点数的最大范围约为 `1.8e308`，可以使用一个常量定义：`math.MaxFloat64`。

打印浮点数时，可以使用`fmt`包配合动词`%f`，代码如下：

```go
package main
import (
        "fmt"
        "math"
)
func main() {
        fmt.Printf("%f\n", math.Pi)
        fmt.Printf("%.2f\n", math.Pi)
}
```

### 3.3、布尔值

Go语言中以`bool`类型进行声明布尔型数据，布尔型数据只有`true（真）`和`false（假）`两个值。

**注意：**

1. 布尔类型变量的默认值为`false`。
2. Go 语言中不允许将整型强制转换为布尔型.
3. 布尔型无法参与数值运算，也无法与其他类型进行转换。

### 3.4、字符串

Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。 Go 语言里的字符串的内部实现使用`UTF-8`编码。 字符串的值为`双引号(")`中的内容，可以在Go语言的源码中直接添加非ASCII码字符，例如：

```go
s1 := "hello"
s2 := "你好"
```

#### 字符串的拼接、占位符

```go
package main
import (
	"fmt"
)
func main() {
	// 字符串拼接
	s1:="hello"
	s2:="world"
	s3:= s1 + s2
	fmt.Println(s3)

	//占位符
	name := "yuan"
	age := 22
	fmt.Printf("姓名：%s，年龄：%d",name,age)


}
```

#### 字符串转义符

Go 语言的字符串常见转义符包含回车、换行、单双引号、制表符等，如下表所示。

| 转义符 |                含义                |
| :----: | :--------------------------------: |
|  `\r`  |         回车符（返回行首）         |
|  `\n`  | 换行符（直接跳到下一行的同列位置） |
|  `\t`  |               制表符               |
|  `\'`  |               单引号               |
|  `\"`  |               双引号               |
|  `\\`  |               反斜杠               |

举个例子，我们要打印一个Windows平台下的一个文件路径：

```go
package main
import (
    "fmt"
)
func main() {
    fmt.Println("str := \"c:\\Code\\lesson01\\go.exe\"")
}
```

#### 多行字符串

Go语言中要定义一个多行字符串时，就必须使用`反引号`字符：

```go
s1 := `第一行
第二行
第三行
`
fmt.Println(s1)
```

反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。

#### 字符串的常用操作

|                方法                 |      介绍      |
| :---------------------------------: | :------------: |
|              len(str)               |     求长度     |
|           +或fmt.Sprintf            |   拼接字符串   |
|            strings.Split            |      分割      |
|          strings.contains           |  判断是否包含  |
| strings.HasPrefix,strings.HasSuffix | 前缀/后缀判断  |
| strings.Index(),strings.LastIndex() | 子串出现的位置 |
| strings.Join(a[]string, sep string) |    join操作    |

### 3.5、类型转换

Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。

强制类型转换的基本语法如下：

```bash
T(表达式)
```

其中，T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等.

```go

f2:= float64(12)
fmt.Println(f2)
```

注意：不是所有转换都可以进行

```go
 f3:= strconv.Itoa(112)
 fmt.Println(f3)
 f4,err := strconv.Atoi("1200")
 fmt.Println(f3)
 fmt.Printf("f4类型：%T " ,f4)
 fmt.Println(err)
```

## 四、运算符

运算符用于在程序运行时执行数学或逻辑运算。

Go 语言内置的运算符有：

- 算术运算符
- 关系运算符
- 逻辑运算符
- 位运算符
- 赋值运算符
- 其他运算符

#### 1、算数运算符

| 运算符 | 描述 |
| :----- | :--- |
| +      | 相加 |
| -      | 相减 |
| *      | 相乘 |
| /      | 相除 |
| %      | 求余 |
| ++     | 自增 |
| --     | 自减 |

 注意：

```go
a := 1
a ++  // 注意：不能写成 ++ a 或 -- a 必须放在右边使用
// b := a++ // 此处为错误的用法，不能写在一行，要单独作为语句使用

fmt.Println(a) // 2
```

#### 2、关系运算符

| 运算符 | 描述                                                         |
| :----- | :----------------------------------------------------------- |
| ==     | 检查两个值是否相等，如果相等返回 True 否则返回 False。       |
| !=     | 检查两个值是否不相等，如果不相等返回 True 否则返回 False。   |
| >      | 检查左边值是否大于右边值，如果是返回 True 否则返回 False。   |
| <      | 检查左边值是否小于右边值，如果是返回 True 否则返回 False。   |
| >=     | 检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。 |
| <=     | 检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。 |

#### 3、逻辑运算符

| 运算符 | 描述                                                         |
| :----- | :----------------------------------------------------------- |
| &&     | 逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。 |
| \|\|   | 逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。 |
| !      | 逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。 |

#### 4、赋值运算符

| 运算符 | 描述                                           |
| :----- | :--------------------------------------------- |
| =      | 简单的赋值运算符，将一个表达式的值赋给一个左值 |
| +=     | 相加后再赋值                                   |
| -=     | 相减后再赋值                                   |
| *=     | 相乘后再赋值                                   |
| /=     | 相除后再赋值                                   |
| %=     | 求余后再赋值                                   |
| <<=    | 左移后赋值                                     |
| >>=    | 右移后赋值                                     |
| &=     | 按位与后赋值                                   |
| ^=     | 按位异或后赋值                                 |
| \|=    | 按位或后赋值                                   |

## 五、流程控制语句

### 3.1 条件语句

条件语句需要开发者通过指定一个或多个条件，并通过测试条件是否为 true 来决定是否执行指定语句，并在条件为 false 的情况在执行另外的语句。

#### 3.1.1、if语句

```go
if 布尔表达式 {
   /* 在布尔表达式为 true 时执行 */
}
```

示例：

```go
// 单分支
	if score > 85{
		fmt.Printf("优秀!")
	}
```

#### 3.1.2、if-else语句

```go
// 双分支
	if score > 85{
		fmt.Printf("优秀!")
	}else {
		fmt.Printf("差！")
	}

```

#### 3.1.3、if-else if-else语句

```go
// 多分支

	if score > 70{
		fmt.Printf("优秀")
	}else if score > 80 {
		fmt.Printf("良")
	}else if score > 60 {
		fmt.Printf("及格")
	}else {
		fmt.Printf("不及格")
	}
	
```

#### 3.1.4、switch语句

```go
switch v1 {
    case v1:
        ...
    case v2:
        ...
    default:
        ...
}
```

示例：

```go
 var weeekday = 3

	switch weeekday {
	case 1: fmt.Printf("星期1")
	case 2: fmt.Printf("星期2")
	case 3: fmt.Printf("星期3")
	case 4: fmt.Printf("星期4")
	case 5: fmt.Printf("星期5")
	default:fmt.Printf("周末！")
	}
```

支持多条件匹配：

```go
switch{
    case 1,2:
    default:
}
```

### 3.2 循环语句

#### 3.2.1、for循环

Go语言中的循环语句只支持for关键字，而不支持while和do-while结构。

方式1:

```go
for init; condition; post { 
}

/*

init： 一般为赋值表达式，给控制变量赋初值;

condition： 关系表达式或逻辑表达式，循环控制条件;

post： 一般为赋值表达式，给控制变量增量或减量。

*/
```

示例：

```go
for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
```

方式2（条件循环）:

关系表达式或逻辑表达式控制循环

```go
for condition { }
```

示例：

```go
count := 0
	for count<10{
	    fmt.Println(count)
	    count ++
	}


/*
无限循环
for true{
	  
	}
*/
```

方式3（遍历循环）:

```go
	// 遍历字符串
    hello := "HelloYuan"
    for k, v := range hello {
		fmt.Printf("%d,%c\n", k, v)
	}
```

#### 3.2.2、嵌套循环

在 for 循环中可以嵌套一个或多个 for 循环，这又分为独立嵌套和关联嵌套。

```
*
**
***
****
*****
******
```

实现上面的三角形：

```go
for i := 1; i < 7; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
```

#### 3.2.3、break和continue

**（1）break**

Go 语言中 break 语句在循环语句中用于以下两方面：

- 用于循环语句中跳出循环，并开始执行循环之后的语句。
- 在多重循环中，可以用标号 label 标出想 break 的循环。

```go
for i := 1; i < 100; i++ {
         if i==14{  // 遇到14退出当前循环
         	break
		 }
		 fmt.Println(":",i)
	}
```

```go
// 九九乘法表
for i := 1; i < 10; i++ {
        /*
		if i>6{
			break
		}
         */

		for j := 1; j <= i; j++ {
            /*
			if j>3{
				break
			}
             */
			fmt.Printf("%d * %d = %2d\t", i, j, i*j) // 放在if上面和下面有没有区别？
		}

		fmt.Println()
	}
```

```go
re:
		for i := 1; i <= 3; i++ {
			fmt.Printf("i: %d\n", i)
			for j := 1; j <= 2; j++ {
				fmt.Printf("j: %d\n", j)
				break re
			}
		}
```

**（2）continue语句**

break用于完全结束一个循环，跳出循环体执行循环后面的语句；而continue是跳过当次循环中剩下的语句，执行下一次循环。简单点说就是break完全结束循环，continue终止本次循环。

```go
for i := 1; i < 100; i++ {
		if i==14{  // 遇到14退出当次循环
			continue
		}
		fmt.Println(":",i)
	}

```

练习作业：

（1）

![](.\assets\QQ截图20201017142436.png)

（2）要求用户输入一个年份和一个月份，判断（要求使用嵌套的if…else和switch分别判断一次）该年该月有多少天。

（3）一只猴子第一天摘下若干个桃子，当即吃了一半，不过瘾，又多吃了一个
第二天早上又将剩下的桃子吃掉一半，又多吃了一个。以后每天早上都吃了前
一天剩下的一半零一个，第10天早上想再吃时，发现只剩下一个桃子了。请问
猴子第一天一共摘了多少个桃子？