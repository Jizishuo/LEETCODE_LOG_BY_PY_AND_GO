package Flip_the_Binary_Tree_226


type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}
	return root
}