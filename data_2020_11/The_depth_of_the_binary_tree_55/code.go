package The_depth_of_the_binary_tree_55

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) +1
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}