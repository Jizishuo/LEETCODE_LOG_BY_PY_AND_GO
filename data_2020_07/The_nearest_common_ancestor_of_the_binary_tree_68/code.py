class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def lowestCommonAncestor(self, root: TreeNode, p: TreeNode, q: TreeNode) -> TreeNode:
        if not root or root==q or root==p:
            return root
        left = self.lowestCommonAncestor(root.left, q, p)
        right = self.lowestCommonAncestor(root.right, q, p)
        if not left and not right:
            return
        if not left:
            return right
        if not right:
            return left
        return root