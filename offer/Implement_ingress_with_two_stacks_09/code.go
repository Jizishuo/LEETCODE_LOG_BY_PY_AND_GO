package Implement_ingress_with_two_stacks_09

//type CQueue struct {
//	in []int
//	out []int
//}
//
//
//func Constructor() CQueue {
//	return CQueue{}
//}
//
//
//func (this *CQueue) AppendTail(value int)  {
//	this.in = append(this.in, value)
//}
//
//
//func (this *CQueue) DeleteHead() int {
//	if len(this.out)==0&&len(this.in)==0 {
//		return -1
//	}
//	if len(this.out) == 0 {
//		for len(this.in) >0{
//			lastindex := len(this.in)-1
//			popv := this.in[lastindex]
//			this.in = this.in[:lastindex]
//			this.out = append(this.out, popv)
//		}
//	}
//	lastidx := len(this.out)-1
//	popv := this.out[lastidx]
//	this.out = this.out[:lastidx]
//	return popv
//}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

type stack []int

func(s *stack) Push(value int) {
	*s = append(*s, value)
}

func(s *stack) Pop() int {
	n := len(*s)
	res := (*s)[n - 1]

	*s = (*s)[:n - 1]

	return res
}

type CQueue struct {
	in stack
	out stack
}


func Constructor() CQueue {
	return CQueue{}
}


func (this *CQueue) AppendTail(value int)  {
	this.in.Push(value)
}


func (this *CQueue) DeleteHead() int {
	if len(this.out) != 0 {
		return this.out.Pop()
	} else if len(this.in) != 0 {
		for len(this.in) != 0 {
			this.out.Push(this.in.Pop())
		}
		return this.out.Pop()
	}
	return -1
}
