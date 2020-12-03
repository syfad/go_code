package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type Shengfen struct {
	ID string
	Name string
	File *os.File
	chanData chan string
}
var wg sync.WaitGroup

//func create_data(ID string, Name string, File *os.File, chanData chan string)  *Shengfen{
//	return &Shengfen{
//		ID,Name,File,chanData,
//	}
//}


func read() {
	pMap := make(map[string]*Shengfen)
	ps := []string{"北京市11", "天津市12", "河北省13",
		"山西省14", "内蒙古自治区15", "辽宁省21", "吉林省22",
		"黑龙江省23", "上海市31", "江苏省32", "浙江省33", "安徽省34",
		"福建省35", "江西省36", "山东省37", "河南省41", "湖北省42",
		"湖南省43", "广东省44", "广西壮族自治区45", "海南省46",
		"重庆市50", "四川省51", "贵州省52", "云南省53", "西藏自治区54",
		"陕西省61", "甘肃省62", "青海省63", "宁夏回族自治区64", "新疆维吾尔自治区65",
		"香港特别行政区81", "澳门特别行政区82", "台湾省83",}

	for _, p := range ps{
		name := (p[:len(p)-2])
		id := (p[len(p)-2:])
		shengfen := Shengfen{ID: id, Name: name}
		pMap[id] = &shengfen
		file, err := os.OpenFile("/Users/sun_admin/Desktop/go7期/data_file/" + shengfen.Name + ".txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		HandleError(err)
		shengfen.File = file
		shengfen.chanData = make(chan string, 1024)
		fmt.Println("管道已创建")
	}
	for _, province := range pMap{
		wg.Add(1)
		go writeFile(province)
		//fmt.Println(province)
	}

	// 读优质文件，写入对应的省份管道
	file, _ := os.Open("./kaifang_good.txt")
	defer file.Close()
	// 缓冲读取
	reader := bufio.NewReader(file)
	// 逐行读
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			for _, province := range pMap {
				close(province.chanData)
				fmt.Println(province.Name, "管道已经关闭")
			}
			break
		}
		// 转str，转utf
		lineStr := string(lineBytes)
		// 逗号切分
		fieldsSlice := strings.Split(lineStr, ",")
		id := fieldsSlice[1][0:2]
		// 对号入座，写入对应管道
		if provice, ok := pMap[id]; ok {
			provice.chanData <- (lineStr + "\n")
		} else {
			fmt.Println("未知的省", id)
		}
	}
}



// 向文件写数据
func writeFile(shenfeng *Shengfen) {
	for lineStr := range shenfeng.chanData {
		shenfeng.File.WriteString(lineStr)
		fmt.Print(shenfeng.Name, "写入", lineStr)
	}
	wg.Done()
}


func HandleError(err error)  {
	if err != nil{
		fmt.Println(err)
	}
}

func main() {
	read()
}