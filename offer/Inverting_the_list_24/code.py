class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        if head == None:
            return head
        p = None
        while head:
            n = head.next
            head.next = p
            p = head
            head = n
        return p