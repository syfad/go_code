package main



import (
"fmt"
"sync"
)

var x = 0

var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	lock.Lock()   // 加锁
	x++
	wg.Done()
	lock.Unlock() // 解锁
}

func main() {

	wg.Add(10000)
	for i:=0;i<10000 ;i++  {
		go add()
	}
	wg.Wait()
	fmt.Println(x)

}