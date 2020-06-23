package main

import (
	"fmt"
	"errors"
)

// select结构体实现通道读
func ReadWithSelect(ch <-chan int) (x int, err error) {
	select {
	case x = <- ch:
		return x, nil
	default:
		return 0, errors.New("channel has no data")
	}
}


// 无缓冲通道读
func ReadNoDataFromBufChWithSelect()  {
	bufCh := make(chan int)
	if v, err := ReadWithSelect(bufCh);err!=nil{
		fmt.Println(err)
	} else {
		fmt.Println("read: %d", v)
	}
	// out put channel has no data
}

// 有缓冲通道读
func ReadNoDataFromNoBufChWithSelect()  {
	bufch := make(chan int, 1)
	if v, err :=ReadWithSelect(bufch);err!=nil{
		fmt.Println(err)
	} else {
		fmt.Println("read: %d", v)
	}
	// out put channel has no data
}
// =======================================================
// select 结构实现通道写
func WriteChWithSelect(ch chan<- int) error {
	select {
	case ch <- 1:
		return nil
	default:
		return errors.New("channel blocked, can not write")
	}
}

func WriteNoBufChWithSelect()  {
	ch := make(chan int)
	if err:=WriteChWithSelect(ch);err!=nil{
		fmt.Println(err)
	} else {
		fmt.Println("write success")
	}
}

// 有缓冲通道写
func WriteBufChButFullWithSelect()  {
	ch := make(chan int, 1)
	// make ch full
	if err := WriteChWithSelect(ch);err!=nil{
		fmt.Println(err)
	} else {
		fmt.Println("write success")
	}
}


func main()  {
	ReadNoDataFromNoBufChWithSelect()
	ReadNoDataFromBufChWithSelect()
	WriteNoBufChWithSelect()
	WriteBufChButFullWithSelect()
}

