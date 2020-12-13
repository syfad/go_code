package main

import (
	"bufio"
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"os"
	"strings"
)

//var reIDcard = `.*[1-9]\d{5}((19\d{2})|(20[0-2]\d))((0[1-9])|(1[0-2]))((0[1-9])|(1\d)|(2\d)|(3[01]))\d{3}[\dXx].*`


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

	goodFile, err1 := os.OpenFile("goodFile.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	HandleError(err1)
	badFile, err1 := os.OpenFile("badFile.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	HandleError(err1)

	defer goodFile.Close()
	defer badFile.Close()

	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF{
			break
		}
		myStr := string(lineBytes)
		lineStr := ConvertEncoding(myStr, "GBK")
		//fmt.Println(myStr)
		//正则匹配，改成匹配len长度
		//re := regexp.MustCompile(reIDcard)
		fields := strings.Split(lineStr, ",")

		if len(fields) >= 2 && len(fields[1]) == 18{
			goodFile.WriteString(lineStr + "\n")
			fmt.Printf("GOOD-File", lineStr)
		}else {
			badFile.WriteString(lineStr + "\n")
			fmt.Printf("BAD-File", lineStr)
		}

		//正则匹配写入的注释了
		//results := re.FindAllStringSubmatch(xxStr, -1)
		////results := re.FindAllStringSubmatch(myStr, -1)
		//for _, res := range results{
		//	//fmt.Println(res[0])
		//
		//	w_lin := bufio.NewWriter(file1)
		//	w_lin.WriteString(res[0]+"\n")
		//	w_lin.Flush()
		//	}
		}
		//fmt.Println(results[1])
	}