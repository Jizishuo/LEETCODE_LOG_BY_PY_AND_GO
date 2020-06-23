package main

import (
	"fmt"
	"sync"
)

func follower(wg *sync.WaitGroup, id int)  {
	fmt.Println("follwer %d find key", id)
	wg.Done()
}

func leader()  {
	var wg sync.WaitGroup
	wg.Add(4)
	for i:=0;i<4;i++ {
		go follower(&wg, i)
	}
	wg.Wait()
	fmt.Println("open teh box together")
}

func main()  {
	leader()
}