package main

import (
	"fmt"
)

func gen() (chan int) {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i:=0;i<3;i++{
			ch <- i
		}

	}()
	return ch
}

func combine(inch1, inch2 <-chan int) <-chan int {
	// 输出管道
	out := make(chan int)
	// 启动协程合并数据
	go func() {
		defer close(out)
		for {
			select {
			case x, open := <- inch1:
				if !open {
					inch1 = nil
					break
				}
				out <- x
			case x, open := <- inch2:
				if !open {
					inch2 = nil
					break
				}
				out <- x
			}
			// 当ch1和ch2都是关闭才退出
			if inch1 == nil && inch2 == nil {
				break
			}
		}
	}()
	return out
}

// break在select内的并不能跳出for-select循环。
func consume(inCh <-chan int) {
	i := 0
	for {
		fmt.Printf("for: %d\n", i)
		select {
		case x, open := <-inCh:
			if !open {
				break
			}
			fmt.Printf("read: %d\n", x)
		}
		i++
	}

	fmt.Println("combine-routine exit")
}

func main()  {
	ch1 := gen()
	ch2 := gen()
	out := combine(ch1, ch2)
	for x:= range out {
		fmt.Println(x)
	}
}
//既然break不能跳出for-select，那怎么办呢？给你3个锦囊：
//在满足条件的case内，使用return，如果有结尾工作，尝试交给defer。
//在select外for内使用break挑出循环，如combine函数。
//使用goto。

// 无阻塞的读、写通道。即使通道是带缓存的，也是存在阻塞的情况，使用select可以完美的解决阻塞读写
// 给某个请求/处理/操作，设置超时时间，一旦超时时间内无法完成，则停止处理。
// select本色：多通道处理