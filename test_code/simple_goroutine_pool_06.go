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