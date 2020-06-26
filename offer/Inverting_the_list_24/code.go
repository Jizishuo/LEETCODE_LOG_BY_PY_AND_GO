package Inverting_the_list_24

type ListNode struct {
	Val int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head==nil {
		return nil
	}
	var p *ListNode
	// c := head
	for head != nil {
		n := head.Next
		head.Next = p
		p = head
		head = n
	}
	return p
}