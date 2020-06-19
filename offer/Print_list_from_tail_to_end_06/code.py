class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def reversePrint(self, head: ListNode) -> List[int]:

        res = []
        while head ==None:
            res.append(head.val)
            head = head.next
        return res[::-1]