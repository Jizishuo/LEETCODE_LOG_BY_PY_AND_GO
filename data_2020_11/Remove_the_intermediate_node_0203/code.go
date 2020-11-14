package Remove_the_intermediate_node_0203

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}