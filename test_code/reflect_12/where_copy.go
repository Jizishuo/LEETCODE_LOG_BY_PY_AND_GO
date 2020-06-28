package main

import (
	"fmt"
)

func main() {
	x := 10
	fmt.Printf("addr of x: %x\n", &x)

	toInterface(x)

	// addr of x: c00000a0f0
	// addr of i: c00003c1f0
}

func toInterface(i interface{}) {
	fmt.Printf("addr of i: %x\n", &i)
}