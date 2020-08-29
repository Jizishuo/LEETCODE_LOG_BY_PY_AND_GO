package Use_the_stack_to_implement_the_queue_232

type MyQueue struct {
	stackpush []int
	stackpop []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue)pushtopop()  {
	if len(this.stackpop) <=0 {
		for _,v := range this.stackpush {
			this.stackpop = append(this.stackpop, v)
		}
		this.stackpush = nil
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.stackpush = append(this.stackpush, x)
	this.pushtopop()
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if (len(this.stackpop)) <= 0 && len(this.stackpush)<=0 {
		return 0
	}
	this.pushtopop()
	ret := this.stackpop[0]
	this.stackpop = this.stackpop[1:]
	return ret
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if (len(this.stackpop)) <=0 && len(this.stackpush)<=0 {
		return 0
	}
	this.pushtopop()
	return this.stackpop[0]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stackpop) <=0 && len(this.stackpush)<=0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */