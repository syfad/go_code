package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func read01(){
	// 文件操作

	file,err:=os.Open("./xiaoshi")
	if err != nil {
		fmt.Println("读取文件异常,",err)
		return
	}

	defer file.Close()
	// 一 读字节
	//var d=make([]byte,1024)
	//n,err:=file.Read(d)
	//fmt.Println(n)
	//fmt.Println(d)
	//fmt.Println(string(d))

	// 读文件
	content:=""
	for true {

		var d=make([]byte,11)
		n,err:=file.Read(d)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("错误",err)
			return
		}
		content+=string(d[:n])
		fmt.Println(content)
	}
	fmt.Printf(content)

}

func read02()  {

	file,err:=os.Open("./xiaoshi")
	if err != nil {
		fmt.Println("读取文件异常,",err)
		return
	}

	defer file.Close()

	reader:=bufio.NewReader(file)
	for true {
		// var line,_,err=reader.ReadLine()
		//fmt.Println(string(line))
		var line,err=reader.ReadString('\n')
		fmt.Printf(line)

		if err == io.EOF {
             break
		}
	}

}

func read3()  {
	content, err := ioutil.ReadFile("xiaoshi")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}






func write01(){

	// 方式1:
	file, err := os.OpenFile("./满江红", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()

	fmt.Println("file",file)

	//str := "满江红\n"
	//file.Write([]byte(str))       //写入字节切片数据
	file.WriteString("xxx") //直接写入字符串数据


}

func write02(){

	file, err := os.OpenFile("满江红2", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	writer.WriteString("大浪淘沙\n") //将数据先写入缓存
	writer.Flush() //将缓存中的内容写入文件

}

func write03()  {
	str := "xxxxx xxxxx"
	err := ioutil.WriteFile("满江红3", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}




func main() {

	read01()

	// write01()
	// write03()
}
