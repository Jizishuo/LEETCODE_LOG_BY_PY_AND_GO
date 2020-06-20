class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def mergeTrees(self, t1: TreeNode, t2: TreeNode) -> TreeNode:
        if t1 is None:
            return t2
        if t2 is None:
            return t1
        tree = TreeNode(t1.val+t2.val)
        tree.left = self.mergeTrees(t1.left, t2.right)
        tree.right = self.mergeTrees(t1.right, t2.right)
        return tree