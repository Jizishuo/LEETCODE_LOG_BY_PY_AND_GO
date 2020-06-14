package main

import (
	"fmt"
	"time"
)

func worker(stopch <- chan struct{})  {
	go func() {
		defer fmt.Println("worker exit..")
		t := time.NewTicker(time.Microsecond*500)

		// 使用停止通道显式退出
		for {
			select {
			case <- stopch:
				fmt.Println("rece stop signal..")
				return
			case <-t.C:
				fmt.Println("working...")
			}
		}
	}()
	return
}

func main()  {
	stopch := make(chan struct{})
	worker(stopch)
	time.Sleep(time.Second*2)
	close(stopch)

	// 等待一下
	time.Sleep(time.Second)
	fmt.Println("main exit...")

}