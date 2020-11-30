package main

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Id int
	RandNum int
}
 type Result struct {
 	job *Job
 	sum int
 }

func main() {
	//1.创建需要的管道
	jobchan := make(chan *Job, 256)
	resultchan := make(chan *Result, 256)

	//2. 穿件workpool
	createpool(256, jobchan, resultchan)


	//3. 开子协程，打印结果
	go func(resultchan chan *Result) {
		for result := range resultchan{
			fmt.Println("job id: %v, randnum: %v, result: %d", result.job.Id, result.job.RandNum, result.sum)
		}
	}(resultchan)

	//4. for 循环，不短创建job，输入任务管道
	var id int
	for {
		id++
		rn := rand.Int()
		job := &Job{
			Id: id,
			RandNum: rn,
		}
		jobchan <- job
	}
}

func createpool(num int, jobchan chan *Job, resultchan chan *Result){
	for i := 0; i < num; i++{
		go func(jobchan chan *Job, resultchan chan *Result) {

		}()
	}
}