package Minimum_stack_151

import "math"

type MinStack struct {
	stack []int
	minstack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minstack: []int{math.MaxInt64},
	}
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)

	this.minstack = append(this.minstack, func(x, y int) int {
												if x>y {
													return y
												} else {
													return x
												}
											}(x, this.minstack[len(this.minstack)-1]))
}


func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack)-1]
	this.minstack = this.minstack[:len(this.minstack)-1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	return this.minstack[len(this.minstack)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */