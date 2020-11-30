package main

import (
	"bufio"
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var reIDcard = `.*[1-9]\d{5}((19\d{2})|(20[0-2]\d))((0[1-9])|(1[0-2]))((0[1-9])|(1\d)|(2\d)|(3[01]))\d{3}[\dXx].*`


var writeChan chan string
var readChan chan string

func main() {
	writeChan = make(chan string)
	read2(writeChan)

	//createPoll(256, writeChan, readChan)
}


func HandleError(err error)  {
	if err != nil{
		fmt.Println(err)
	}
}
func read()  {
	conntentBytes, err := ioutil.ReadFile("/Users/sun_admin/Desktop/go7期/数据/省份.txt")
	HandleError(err)
	conntenStr := string(conntentBytes)
	lineStrs := strings.Split(conntenStr, "\n\r")
	for _, lineStr := range lineStrs{
		fmt.Println(lineStr)
		newStr := ConvertEncoding(lineStr, "GBk")
		fmt.Println(newStr)
	}

}

func ConvertEncoding(srcStr string, encoding string)(dstStr string){
	enc := mahonia.NewDecoder(encoding)
	dstStr = enc.ConvertString(srcStr)
	return
}

func read2(c chan string)  {
	//file, err := os.Open("/Users/sun_admin/Desktop/go7期/数据/kaifang.txt")
	file, err := os.Open("/Users/sun_admin/Desktop/go7期/shot.txt")
	defer file.Close()
	HandleError(err)
	reader := bufio.NewReader(file)
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF{
			break
		}
		myStr := string(lineBytes)
		//xxStr := ConvertEncoding(myStr, "GBK")
		//fmt.Println(myStr)
		re := regexp.MustCompile(reIDcard)
		results := re.FindAllStringSubmatch(myStr, -1)
		for _, res := range results{
			//fmt.Println(res[0])
			c <- res[0]
		}
		//fmt.Println(results[1])
	}

}