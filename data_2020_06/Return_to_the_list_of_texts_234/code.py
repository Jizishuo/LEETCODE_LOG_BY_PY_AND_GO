class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def isPalindrome(self, head: ListNode) -> bool:

        # s, f = head, head
        # p = None
        # while f!=None and f.next!=None:
        #     f = f.next.next
        #
        #     t = s
        #     s = s.next
        #     t.next = p
        #     p = t
        #
        # if f!=None:
        #     s = s.next
        #
        # while s!=None and p != None:
        #     if s.val!= p.val:
        #         return False
        #     s= s.next
        #     p = p.next
        # return True
        s = []
        c = head
        while c:
            s.append(c)
            c =c.next

        n = head
        while s:
            n2 = s.pop()
            if n.val != n2.val:
                return False
            n = n.next
        return True