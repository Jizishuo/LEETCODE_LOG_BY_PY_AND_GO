class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def hasCycle(self, head: ListNode) -> bool:
        f, s = head

        while f and f.next:
            s = s.next
            f = f.next.next
            if s is f:
                return True
        return False