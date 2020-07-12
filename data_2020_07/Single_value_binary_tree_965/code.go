package Single_value_binary_tree_965

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if (root.Left !=nil && root.Left.Val !=root.Val) && (root.Right !=nil && root.Right.Val !=root.Val) {
		return false
	}
	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}