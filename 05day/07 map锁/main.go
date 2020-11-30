package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	// m := make(map[int]int)

	m:=sync.Map{}   // 自带锁的map



	// 添一些假数据
	//for i := 0; i < 5; i++ {
	//	m[i] = i*i
	//}  // map[0:0 1:1 2:4 3:9 4:16]
	// 遍历打印
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(i int ) {
			// fmt.Println(m[i])
			// m[i] = i*i
			m.Store(i,i*i)
			wg.Done()
		}(i)    // 思考  i 不传入会怎样

	}
	wg.Wait()
	fmt.Println(m)
	fmt.Println(m.Load(3))
	fmt.Println(m.Load(4))

}
