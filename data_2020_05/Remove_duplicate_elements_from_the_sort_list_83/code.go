package Remove_duplicate_elements_from_the_sort_list_83



// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	//node := head
	//for node != nil {
	//	val := node.Val
	//	for node.Next != nil && node.Next.Val == val {
	//		node.Next = node.Next.Next
	//	}
	//	node = node.Next
	//}
	//return head
	if head == nil || head.Next == nil {
		return head
	}
	head.Next = deleteDuplicates(head.Next)
	if head.Val == head.Next.Val {
		head = head.Next
	}
	return head
}