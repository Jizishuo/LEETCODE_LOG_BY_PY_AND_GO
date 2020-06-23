package main

import (
	"fmt"
	"errors"
	"time"
)

// select结构体实现通道读
func ReadWithSelect1(ch <-chan int) (x int, err error) {
	timeout := time.NewTimer(time.Microsecond*500)
	select {
	case x = <- ch:
		return x, nil
	case <- timeout.C:
		return 0, errors.New("channel has no data")
	}
}


// 无缓冲通道读
func ReadNoDataFromBufChWithSelect1()  {
	bufCh := make(chan int)
	if v, err := ReadWithSelect1(bufCh);err!=nil{
		fmt.Println(err)
	} else {
		fmt.Println("read: %d", v)
	}
	// out put channel has no data
}

// 有缓冲通道读
func ReadNoDataFromNoBufChWithSelect1()  {
	bufch := make(chan int, 1)
	if v, err :=ReadWithSelect1(bufch);err!=nil{
		fmt.Println(err)
	} else {
		fmt.Println("read: %d", v)
	}
	// out put channel has no data
}
// =======================================================
// select 结构实现通道写
func WriteChWithSelect1(ch chan<- int) error {
	timeout := time.NewTimer(time.Microsecond * 500)

	select {
	case ch <- 1:
		return nil
	case <- timeout.C:
		return errors.New("channel blocked, can not write")
	}
}

func WriteNoBufChWithSelect1()  {
	ch := make(chan int)
	if err:=WriteChWithSelect1(ch);err!=nil{
		fmt.Println(err)
	} else {
		fmt.Println("write success")
	}
}

// 有缓冲通道写
func WriteBufChButFullWithSelect1()  {
	ch := make(chan int, 1)
	// make ch full
	if err := WriteChWithSelect1(ch);err!=nil{
		fmt.Println(err)
	} else {
		fmt.Println("write success")
	}
}


func main()  {
	ReadNoDataFromNoBufChWithSelect1()
	ReadNoDataFromBufChWithSelect1()
	WriteNoBufChWithSelect1()
	WriteBufChButFullWithSelect1()
}



