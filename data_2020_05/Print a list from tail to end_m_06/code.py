class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def reversePrint(self, head: ListNode) -> List[int]:
        if head == None:
            return head
        l = head
        r = head.next
        while r != None:
            tt = r.next
            r.next = l
            l = r
            r = tt
        head.next = None

        res = []
        while l != None:
            res.append(l.val)
        return res