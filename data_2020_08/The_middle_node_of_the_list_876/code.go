package The_middle_node_of_the_list_876

type ListNode struct {
    Val int
    Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	//f:= head
	//s := head
	//
	//for f!=nil && f.Next!=nil {
	//	s = s.Next
	//	f= f.Next.Next
	//}
	//return s
	var ll []*ListNode
	for head!=nil {
		ll = append(ll, head)
		head = head.Next
	}
	return ll[len(ll)/2]
}