package main

import (
	"fmt"
	"time"
)

func producer(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer func() {
			close(out)
			out = nil
			fmt.Println("producer exit...")
		}()

		for i:=0;i<n;i++ {
			fmt.Printf("send %d\n", i)
			out <- i
			time.Sleep(time.Microsecond)
		}
	}()
	return out
}

func consumer(in <-chan int) <-chan struct{} {
	finish := make(chan struct{})

	go func() {
		defer func() {
			fmt.Println("working exit..")
			finish <- struct{}{}
			close(finish)
		}()

		// 使用 for 范围退出 go 例程
		// 范围能够检测通道的关闭/结束
		for x := range in {
			fmt.Println("process %d\n", x)
		}
	}()
	return finish
}

func main()  {
	out := producer(3)
	finish := consumer(out)

	// 等待 consumer 退出
	<- finish
	fmt.Println("main exit...")
}

//// send 0
//// Process 0
//// send 1
//// Process 1
//// send 2
//// Process 2
//// producer exit
//// worker exit
//// main exit