package main

import (
	"fmt"
	"sync"
	"time"
)

// 效率对比

// 声明读写锁
var rwlock sync.RWMutex
var mutex sync.Mutex
var wg sync.WaitGroup
// 全局变量
var x int

// 写数据
func write() {

	// mutex.Lock()
	rwlock.Lock()
	x += 1
	time.Sleep(10 * time.Millisecond)
	// mutex.Unlock()
	rwlock.Unlock()


	wg.Done()
}

func read(i int) {

	// mutex.Lock()
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	fmt.Println(x)
	// mutex.Unlock()
    rwlock.RUnlock()

	wg.Done()
}


func main() {
	start := time.Now()
	wg.Add(1)
	go write()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read(i)
	}

	wg.Wait()

	fmt.Println("运行时间：", time.Now().Sub(start))

}