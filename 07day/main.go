package main

import (
	"fmt"
	"github.com/axgle/mahonia"
	"os"
)

var reIDcard = `.*[1-9]\d{5}((19\d{2})|(20[0-2]\d))((0[1-9])|(1[0-2]))((0[1-9])|(1\d)|(2\d)|(3[01]))\d{3}[\dXx].*`


func main() {

	ps := []string{"北京市11", "天津市12", "河北省13", "山西省14", "内蒙古自治区15", "辽宁省21", "吉林省22",}
	fmt.Printf("%T", ps)
}


func HandleError(err error)  {
	if err != nil{
		fmt.Println(err)
	}
}


func ConvertEncoding(srcStr string, encoding string)(dstStr string){
	enc := mahonia.NewDecoder(encoding)
	dstStr = enc.ConvertString(srcStr)
	return
}

func read1() {
	file, err := os.Open("/Users/sun_admin/Desktop/go7期/数据/kaifang.txt")
	HandleError(err)
	defer file.Close()
}
