# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:

    # def __init__(self):
    #     self.vv = None
    def isUnivalTree(self, root: TreeNode) -> bool:
        if root == None:
            return True

        # if not self.vv:
        #     self.vv = root.val
        # if root.val != self.vv:
        #     return False
        if (root.left != None and root.left.val !=root.val) or (root.right != None and root.right.val !=root.val):
            return False

        return self.isUnivalTree(root.left) and self.isUnivalTree(root.right)
