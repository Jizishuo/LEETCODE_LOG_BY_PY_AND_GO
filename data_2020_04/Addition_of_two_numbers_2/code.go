package Addition_of_two_numbers_2

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	sum := l1.Val + l2.Val
	nextNode := addTwoNumbers(l1.Next, l2.Next)
	if sum < 10 {
		return &ListNode{sum,nextNode}
	} else {
		tempNode := &ListNode{
			Val:  1,
			Next: nil,
		}
		return &ListNode{
			Val:  sum-10,
			Next: addTwoNumbers(nextNode, tempNode),
		}
	}


}