package Interstated_list_160

type ListNode struct {
    Val int
    Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    a , b := headA, headB

    for a != b {
        if a != nil {
            a =a.Next
        } else {
            a = headB
        }
        if b != nil {
            b = b.Next
        }else {
            b = headA
        }
    }
    return a
}