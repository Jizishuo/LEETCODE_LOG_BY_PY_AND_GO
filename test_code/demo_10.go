package main

import (
	"sync"
	"time"
	"math/rand"
	"fmt"
)

func handle(wg *sync.WaitGroup, a int) chan int {
	out := make(chan int)
	go func() {
		time.Sleep(time.Duration(rand.Intn(3))*time.Second)
		out <- a
		wg.Done()
	}()
	return out
}

func main()  {
	reqs := []int{1,2,3,4,5,6,7,8,9}

	//存放结果的chan的chan
	outs :=make(chan chan int, len(reqs))
	var wg sync.WaitGroup
	wg.Add(len(reqs))

	for _, x := range reqs {
		o := handle(&wg, x)
		outs <- o
	}
	go func() {
		wg.Wait()
		close(outs)
	}()

	// 读取结果，结果有序
	for o := range outs {
		fmt.Println(<- o)
	}
}