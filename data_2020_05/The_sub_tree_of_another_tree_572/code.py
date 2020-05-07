# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def isSubtree(self, s: TreeNode, t: TreeNode) -> bool:
        if not s and not t:
            return True
        if not s or not t:
            return False

        return self.issame(s, t) or self.isSubtree(s.left, t) or self.isSubtree(s.right, t)

    def issame(self, s, t):
        if not s and not t:
            return True
        if not s or not t:
            return False
        return (s.val==t.val) and self.issame(s.left, t.left) and self.issame(s.right, t.right)