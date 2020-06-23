class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def getKthFromEnd(self, head: ListNode, k: int) -> ListNode:

        f, s = head, head
        for _ in range(k):
            f = f.next

        while f:
            f, s = f.next, s.next
        return s