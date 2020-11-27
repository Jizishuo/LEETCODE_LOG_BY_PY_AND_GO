package Returns_the_penultimate_node_0202

type ListNode struct {
	Val  int
	Next *ListNode
}

func kthToLast(head *ListNode, k int) int {
	f := head
	s := head
	for k>0 {
		f = f.Next
		k--
	}

	for f != nil {
		f = f.Next
		s = s.Next
	}
	return s.Val
}