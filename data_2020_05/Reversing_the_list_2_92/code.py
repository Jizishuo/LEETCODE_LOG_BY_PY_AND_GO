class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def reverseBetween(self, head: ListNode, m: int, n: int) -> ListNode:
        if m == n:
            return head
        demmy = ListNode(-1)
        demmy.next = head

        a, d = demmy, demmy
        for _ in range(m-1):
            a = a.next

        for _ in range(n):
            d = d.next

        b, c = a.next, d.next

        per = b
        cur = per.next
        while cur != c:
            next = cur.next
            cur.next = per
            per = cur
            cur = next
        a.next = d
        b.next = c
        return demmy.next
