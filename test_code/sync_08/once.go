package main

import (
	"sync"
	"fmt"
)

func main()  {
	var once sync.Once
	oncebody := func() {
		fmt.Println("only once..")
	}

	done := make(chan bool)

	for i:=0;i<10;i++ {
		go func() {
			once.Do(oncebody)
			done <- true
		}()
	}

	for i:=0;i<10;i++ {
		<- done
	}
}