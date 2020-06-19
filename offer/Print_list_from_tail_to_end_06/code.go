package Print_list_from_tail_to_end_06

type ListNode struct {
    Val int
    Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var res []int

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	for i, j := 0, len(res)-1; i<j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res


}