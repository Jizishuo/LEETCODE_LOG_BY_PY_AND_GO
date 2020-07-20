package Triad_0301

type TripleInOne struct {
	d []int
}


func Constructor(stackSize int) TripleInOne {
	d := make([]int, stackSize*3+4)
	copy(d, []int{4,5,6, stackSize*3+3})
	return TripleInOne{d}
}


func (this *TripleInOne) Push(stackNum int, value int)  {
	if this.d[stackNum] <= this.d[3] {
		this.d[this.d[stackNum]] = value
		this.d[stackNum] +=3
	}
}


func (this *TripleInOne) Pop(stackNum int) int {
	if this.d[stackNum] >6 {
		this.d[stackNum] -= 3
		return this.d[this.d[stackNum]]
	} else {
		return -1
	}
}


func (this *TripleInOne) Peek(stackNum int) int {
	if this.d[stackNum] <= 6 {
		return -1
	}
	return this.d[this.d[stackNum]-3]
}


func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return this.d[stackNum] <= 6
}


/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */