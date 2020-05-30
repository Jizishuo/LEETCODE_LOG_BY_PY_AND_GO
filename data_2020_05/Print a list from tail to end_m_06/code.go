package Print_a_list_from_tail_to_end_m_06


type ListNode struct {
	Val int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	//res := []int{}
	//for head.Next != nil {
	//	res = append(res, head.Next.Val)
	//	head= head.Next
	//}
	//res2 := []int{}
	//for i := len(res);i>0; i-- {
	//	res = append(res2, res[i])
	//}
	//return res2

	var n *ListNode

	for head != nil {
		head, head.Next, n = head.Next, n, head
	}
	res := []int{}
	for n != nil {
		res = append(res, n.Val)
		n = n.Next
	}

	return res
}