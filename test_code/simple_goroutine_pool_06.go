package main

import (
	"fmt"
	"time"
)

func genjob(n int) <-chan int {
	jobch := make(chan int, 200)
	go func() {
		for i:=0;i<n;i++ {
			jobch <- i
		}
		close(jobch)
	}()
	return jobch
}

func worker(id int, jobch <-chan int, retch chan<- string)  {
	cnt := 0
	for job := range jobch {
		cnt ++
		ret := fmt.Sprintf("worker %d processed job: %d, it's the %dth processed by me.", id, job, cnt)
		retch <- ret
	}
}

func workerpool(n int, jobch <-chan int, retch chan<- string)  {
	for i:=0;i<n;i++ {
		go worker(i, jobch, retch)
	}
}

func main()  {
	jobch := genjob(10000)
	retch := make(chan string, 10000)
	workerpool(5, jobch, retch)

	time.Sleep(time.Second)
	close(retch)
	for ret := range retch {
		fmt.Println(ret)
	}
}

// 比如，我们有A、B两类工作，不想把太多资源花费在B类务上，而是花在A类任务上。对于A，我们可以来1个开一个goroutine去处理，对于B，我们可以使用一个协程池，协程池里有5个线程去处理B类任务，这样B消耗的资源就不会太多。
