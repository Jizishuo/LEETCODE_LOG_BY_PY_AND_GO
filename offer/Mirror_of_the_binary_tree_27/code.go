package Mirror_of_the_binary_tree_27

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root != nil{
		root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
	}
	return root
}