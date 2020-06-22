package Remove_nodes_from_the_list_18

type ListNode struct {
    Val int
    Next *ListNode
}


func deleteNode(head *ListNode, val int) *ListNode {
	n := &ListNode{0,head}
	f := n
	s := n.Next

	for {
		if s.Val == val {
			f.Next = s.Next
			break
		}
		f = f.Next
		s = s.Next
	}
	return n.Next
}
