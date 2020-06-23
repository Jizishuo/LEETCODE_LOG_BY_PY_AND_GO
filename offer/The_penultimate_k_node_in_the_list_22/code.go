package The_penultimate_k_node_in_the_list_22


type ListNode struct {
    Val int
    Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	f, s := head, head
	for i:=0;i<k;i++ {
		f = f.Next
	}

	for f != nil {
		f, s = f.Next, s.Next
	}
	return s
}