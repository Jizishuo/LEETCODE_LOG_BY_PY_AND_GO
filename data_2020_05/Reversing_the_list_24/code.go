package Reversing_the_list_24

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	//if head==nil {
	//	return head
	//}
	//var newl *ListNode
	//for head != nil {
	//	head, head.Next, newl = head.Next, newl, head
	//}
	//return newl
	l := head
	r := head.Next
	for r!=nil {
		tmp := r.Next
		r.Next = l
		l = r
		r = tmp
	}
	head.Next = nil
	return l
}