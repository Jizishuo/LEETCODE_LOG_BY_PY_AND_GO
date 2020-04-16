package Using_queue_to_implement_stack_225

type MyStack struct {
	q []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.q = append(this.q, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if len(this.q) == 0 {
		return -1
	}
	r := this.q[len(this.q)-1]
	this.q = this.q[:len(this.q)-1]
	return r
}


/** Get the top element. */
func (this *MyStack) Top() int {
	if len(this.q) == 0 {
		return -1
	}
	return this.q[len(this.q)-1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.q)==0
}