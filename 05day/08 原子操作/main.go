package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 效率对比
// 原子操作需要接收int32或int64
var x int32
var wg sync.WaitGroup
var mutex sync.Mutex

// 互斥锁操作
func add1() {
	for i := 0; i < 500; i++ {
		mutex.Lock()
		x += 1
		mutex.Unlock()
	}
	wg.Done()
}

// 原子操作
func add2() {
	for i := 0; i < 500; i++ {
		atomic.AddInt32(&x, 1)
	}
	wg.Done()
}


func main() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		//go add1()
		go add2()
	}
	wg.Wait()
	fmt.Println("x:", x)
	fmt.Println("执行时间：", time.Now().Sub(start))
}