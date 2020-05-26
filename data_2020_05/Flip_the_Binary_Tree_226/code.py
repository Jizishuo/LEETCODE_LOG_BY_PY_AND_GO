class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def invertTree(self, root: TreeNode) -> TreeNode:

        if root != None:
            root.left, root.right = self.invertTree(root.right), self.invertTree(root.left)
        return root