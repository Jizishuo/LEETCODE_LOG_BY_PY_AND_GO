//package Numbers_that_only_appear_once_136

package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		fmt.Print(v)
		res = v ^ res
		fmt.Print(res)
		fmt.Printf("********************\n")
	}
	return res
}

func main () {
	ddd := []int{1, 2, 34, 33}
	// ddd := [2,2,1]
	singleNumber(ddd)
}