class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def middleNode(self, head: ListNode) -> ListNode:
        if head.next==None:
            return head
        f = head
        s = head
        while f and f.next:
            s = s.next
            f = f.next.next
        return s
