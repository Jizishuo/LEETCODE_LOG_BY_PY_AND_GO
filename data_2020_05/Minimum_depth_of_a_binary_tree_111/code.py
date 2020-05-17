class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:


    def minDepth(self, root: TreeNode) -> int:
        if root == None:
            return 0
        queue = [(1, root)]

        while queue:
            d, n = queue.pop(0)
            if not n.left and not n.right:
                return d
            if n.left:
                queue.append((d+1, n.left))
            if n.right:
                queue.append((d+1, n.right))