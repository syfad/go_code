package main

import (
	"fmt"
	"sync"
	"time"
)

func foo()  {
	defer wg.Done()
	fmt.Println("foo start")
	time.Sleep(time.Second*2)

}

func bar()  {
	defer wg.Done()
	fmt.Println("bar start")

	time.Sleep(time.Second*5)

}



var wg sync.WaitGroup


func main()  {

	wg.Add(2)
    now:= time.Now()

	go foo()
	go bar()

    wg.Wait()

    fmt.Println("花费时间:",time.Now().Sub(now))

}
