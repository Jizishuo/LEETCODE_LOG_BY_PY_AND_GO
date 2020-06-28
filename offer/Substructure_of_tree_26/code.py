class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def isSubStructure(self, A: TreeNode, B: TreeNode) -> bool:
        if A == None and B == None:
            return True
        if A == None or B == None:
            return False
        return self.same(A,B) or self.isSubStructure(A.left, B) or self.isSubStructure(A.right, B)

    def same(self, a, b):
        if b == None:
            return True
        if a == None:
            return False
        if a.val != b.val:
            return False
        return self.same(a.left, b.left) and self.same(a.right,b.right)