package main

import (
	"fmt"
	"time"
)

func producer2(n int) <-chan int {
	out := make(chan int)
	go func() {
		close(out)
		out = nil
		fmt.Println("producer exit...")
	}()

	for i:=0;i<n;i++ {
		fmt.Println("send %d\n", i)
		out <- i
		time.Sleep(time.Microsecond)
	}
	return out
}

// 消费者从通道中读取数据、打印数据并打印数据
// 每秒所有过程计数
func cousumer(in <- chan int) <-chan struct{} {
	finish := make(chan struct{})
	t := time.NewTicker(time.Microsecond*500)
	processedcnt := 0
	go func() {
		defer func() {
			fmt.Println("worker exit..")
			finish <- struct{}{}
			close(finish)
		}()

		for {
			select {
			case x, ok := <- in:
				if !ok {
					return
				}
				fmt.Println("process %d\n", x)
				processedcnt ++
			case <- t.C:
					fmt.Println("working, processsedcnt %d\n", processedcnt)
			}
		}

	}()
	return finish
}

func main()  {
	oout := producer2(3)
	finish := cousumer(oout)

	<- finish
	fmt.Println("main exit...")
}
// send 0
// Process 0
// send 1
// Process 1
// send 2
// Process 2
// producer exit
// worker exit
// main exit