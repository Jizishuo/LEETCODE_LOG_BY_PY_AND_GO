package Remove_list_elements_203

type ListNode struct {
    Val int
    Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	p := &ListNode{}
	n_l := p
	n_l.Next = head
	for head != nil {
		if head.Val == val {
			n_l.Next=head.Next
		} else {
			n_l = head
		}
		head = head.Next
	}
	return p.Next

	var dummyNode = &ListNode{}
	var pre = dummyNode
	pre.Next= head

	for head != nil{
		if head.Val == val{
			pre.Next = head.Next
		}else {
			pre = head
		}
		head = head.Next
	}

	return dummyNode.Next


}