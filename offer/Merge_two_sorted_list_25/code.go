package Merge_two_sorted_list_25

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := &ListNode{}
	res :=tmp
	for l1!=nil && l2 !=nil {
		if l1.Val<l2.Val {
			tmp.Next=l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 == nil {
		tmp.Next = l2
	} else {
		tmp.Next = l1
	}
	return res.Next

	// 	tmp := &ListNode{}
	//	res := tmp
	//	for l1 != nil && l2 != nil {
	//		if l1.Val < l2.Val {
	//			tmp.Next = l1
	//			l1 = l1.Next
	//		} else {
	//			tmp.Next = l2
	//			l2 = l2.Next
	//		}
	//		tmp = tmp.Next
	//	}
	//	if l1 == nil {
	//		tmp.Next = l2
	//	} else {
	//		tmp.Next = l1
	//	}
	//	return res.Next
}
