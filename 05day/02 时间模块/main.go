package main

import (
	"fmt"
	"time"
)

func main()  {



  // now := time.Now()
  //fmt.Println(now)
  //
  //fmt.Println(now.Year())
  //fmt.Println(now.Day())
  //
  //
  //fmt.Println(now.Unix())
  //fmt.Println(now.UnixNano())
  //
  //
  //time.Sleep(time.Second*3)


  //later:=now.Add(time.Hour*24*3)
  //fmt.Println(later)

   // fmt.Println(time.Now().Sub(now))


	// 格式输出
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))


}
