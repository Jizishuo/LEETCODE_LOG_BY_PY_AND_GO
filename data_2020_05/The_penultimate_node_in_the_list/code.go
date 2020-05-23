package The_penultimate_node_in_the_list


type ListNode struct {
	Val int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	//var res []*ListNode
	//
	//for head != nil {
	//	res = append(res, head)
	//	head = head.Next
	//}
	//l := len(res)
	//if l>=k {
	//	return res[l-k]
	//}
	//return nil

	f, m := head, head

	for i:=0;i<k;i++ {
		f = f.Next
	}

	for f != nil {
		f,m = f.Next, m.Next
	}
	return m
}

