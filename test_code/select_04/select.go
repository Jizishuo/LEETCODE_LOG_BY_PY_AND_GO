package select_04

import (
	"fmt"
	"math/rand"
	"time"
)

func eat() chan string {
	out := make(chan string)
	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(5))*time.Second)
		out <- "mom call you eating"
		close(out)
	}()
	return out
}

func main()  {
	eatCh := eat()
	sleep := time.NewTimer(time.Second*3)
	select {
	case s := <- eatCh:
		fmt.Println(s)
	case <- sleep.C:
		fmt.Println("time to sleep")
	default:
		fmt.Println("beat doudou..")
	}
}