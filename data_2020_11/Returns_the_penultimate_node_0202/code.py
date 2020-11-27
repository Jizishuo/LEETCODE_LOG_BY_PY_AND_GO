class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def kthToLast(self, head: ListNode, k: int) -> int:

        f = head
        s = head
        for _ in range(k):
            f = f.next

        while f:
            s = s.next
            f = f.next
        return s.val