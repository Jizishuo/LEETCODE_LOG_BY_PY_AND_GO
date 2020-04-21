package Maximum_depth_of_a_binary_tree_104

/**
* Definition for a binary tree node.

*/
type TreeNode struct {

    Val int
    Left *TreeNode
    Right *TreeNode
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		return func (x, y int) int {
			if x>y {
				return x
			} else {
				return y
			}
		}(maxDepth(root.Left)+1, maxDepth(root.Right)+1)
	}
}
