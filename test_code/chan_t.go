package main

import (
	"fmt"
)

func producer(nums ...int) <- chan int  {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}
func square(inCh <-chan int) <- chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range inCh {
			out <- n * n
		}
	}()
	return out
}

func main()  {
	in := producer(1,2,3,4)
	ch := square(in)
	for ret := range ch {
		fmt.Printf("%3d\n", ret)
	}
	fmt.Println("----")
}