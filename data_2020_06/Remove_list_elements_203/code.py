class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def removeElements(self, head: ListNode, val: int) -> ListNode:
        newH = ListNode(0)
        newH.next = head
        pre, p = newH, head
        while p:
            if p.val == val:
                pre.next = p = p.next
            else:
                pre, p = p, p.next
        return newH.next



