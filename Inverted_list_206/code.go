package Inverted_list_206

func reverseList(head *ListNode) *ListNode {
	//if head == nil || head.Next == nil {
	//	return head
	//}
	//p := reverseList(head.Next)
	//head.Next.Next = head
	//head.Next = nil
	//return p
	//

	var newHead, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = newHead
		newHead = head
		head = next
	}
	return newHead
}