package Return_to_the_list_of_texts_234

type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head ==nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	var per *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		tmp := slow
		slow = slow.Next
		tmp.Next = per
		per = tmp
	}
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil && per != nil {
		if slow.Val != per.Val {
			return false
		}
		slow = slow.Next
		per = per.Next
	}
	return true
}
//func hhh(l *ListNode) *ListNode {
//	if l.Next == nil {
//		return l
//	}
//}