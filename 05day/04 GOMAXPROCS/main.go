package main

import (
	"fmt"
	"runtime"
	"sync"
)


var wg sync.WaitGroup

func foo() {
	for i := 1; i < 1000; i++ {
		fmt.Println("A:", i)


	}

	wg.Done()
}

func bar() {
	for i := 1; i < 1000; i++ {
		fmt.Println("B:", i)
		// time.Sleep(time.Millisecond*30)
		// time.Sleep(time.Millisecond*20)
		if (i==100){
			// time.Sleep(time.Millisecond*20)
			runtime.Gosched()  //  调度函数
		}
	}

	wg.Done()
}




func main()  {

	// 打印CPU数量
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
    wg.Add(2)
	go foo()
	go bar()

	wg.Wait()




}
