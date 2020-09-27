package Print_in_order_1114

import (
	"fmt"
)

var secendchan = make(chan int)
var thirdchan = make(chan int)
var mainchan = make(chan int)

func first() {
	fmt.Printf("first")
	secendchan <-1
}

func secend() {
	siganl := <- secendchan
	fmt.Printf("secend")
	thirdchan <- siganl
	close(secendchan)
}

func third() {
	signal := <- thirdchan
	fmt.Printf("third")
	mainchan <- signal
	close(thirdchan)
}

func main()  {
	funcmap := map[int]func(){1:first, 2:secend, 3:third}
	inputlist :=[3]int{1,2,3}
	for _, num := range inputlist {
		go funcmap[num]()
	}
	_ = <- mainchan
	close(mainchan)

}