package Balance_the_binary_tree_110

import "math"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return root == nil || isBalanced(root.Left) &&
		math.Abs(height(root.Left)-height(root.Right))< 2 &&
		isBalanced(root.Right)
}

func height(node *TreeNode) float64 {
	if node == nil {
		return 0
	}
	return math.Max(height(node.Left), height(node.Right))+1
}