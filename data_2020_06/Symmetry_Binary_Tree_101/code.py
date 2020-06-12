class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def isSymmetric(self, root: TreeNode) -> bool:
        def dfs(root1, root2):
            if root1 == root2: return True
            if not root1 or not root2: return False
            if root1.val != root2.val: return False
            return dfs(root1.left, root2.right) and dfs(root1.right, root2.left)

        if not root: return True
        return dfs(root.left, root.right)
