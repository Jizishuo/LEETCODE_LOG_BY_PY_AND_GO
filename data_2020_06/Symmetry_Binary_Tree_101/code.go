package Symmetry_Binary_Tree_101

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirrorRecursive(root.Left, root.Right)
}

func isMirrorRecursive(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && isMirrorRecursive(left.Left, right.Right) && isMirrorRecursive(left.Right, right.Left)
}