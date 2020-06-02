class Solution:
    def isBalanced(self, root: TreeNode) -> bool:
        return root==None or self.isBalanced(root.left) and abs(self.hhh(root.left)-self.hhh(root.right))<2 and self.isBalanced(root.right)
    def hhh(self, root):
        if root == nil:
            return 0

        return max(self.hhh(root.left), self.hhh(root.right))+1