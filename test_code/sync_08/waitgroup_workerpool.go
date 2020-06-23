package main

import (
	"fmt"
	"sync"
)

func getjob(n int) <-chan int {
	jobch := make(chan int, 200)
	go func() {
		for i:=0;i<n;i++ {
			jobch <- i
		}
		close(jobch)
	}()
	return jobch
}

func worker(wg *sync.WaitGroup, id int, jobch <-chan int, retch chan<- string)  {
	cnt := 0
	for job := range jobch {
		cnt ++
		res := fmt.Sprintf("worker %d processed job: %d, it's the %dth processed by me.", id, job, cnt)
		retch <- res
	}
	wg.Done()
}

func workerpool(n int, jobch <-chan int, retch chan<- string)  {
	var wg sync.WaitGroup
	wg.Add(n)
	for i:=0;i<n;i++ {
		go worker(&wg, i, jobch, retch)
	}
	wg.Wait()
	close(retch)
}


func main()  {
	jobch := getjob(10)
	retch := make(chan string, 100000)
	workerpool(5, jobch, retch)
	for ret := range retch {
		fmt.Println(ret)
	}
}