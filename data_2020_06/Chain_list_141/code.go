package Chain_list_141

type ListNode struct {
    Val int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	f := head.Next
	for f !=nil && head !=nil && f.Next !=nil  {
		if f == head {
			return true
		}
		f = f.Next.Next
		head = head.Next
	}
	return false
}