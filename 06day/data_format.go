package main

import (
	"bufio"
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"os"
	"regexp"
)

var reIDcard = `.*[1-9]\d{5}((19\d{2})|(20[0-2]\d))((0[1-9])|(1[0-2]))((0[1-9])|(1\d)|(2\d)|(3[01]))\d{3}[\dXx].*`


func main() {
	read2()
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

func read2()  {
	file, err := os.Open("/Users/sun_admin/Desktop/go7期/数据/kaifang.txt")
	HandleError(err)
	//file, err := os.Open("/Users/sun_admin/Desktop/go7期/shot.txt")
	defer file.Close()
	reader := bufio.NewReader(file)

	file1, err1 := os.OpenFile("format_after.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	HandleError(err1)
	defer file1.Close()

	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF{
			break
		}
		myStr := string(lineBytes)
		xxStr := ConvertEncoding(myStr, "GBK")
		//fmt.Println(myStr)
		re := regexp.MustCompile(reIDcard)
		results := re.FindAllStringSubmatch(xxStr, -1)
		//results := re.FindAllStringSubmatch(myStr, -1)
		for _, res := range results{
			//fmt.Println(res[0])

			w_lin := bufio.NewWriter(file1)
			w_lin.WriteString(res[0]+"\n")
			w_lin.Flush()
			}
		}
		//fmt.Println(results[1])
	}