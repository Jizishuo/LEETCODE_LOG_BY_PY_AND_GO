package The_diameter_of_the_binary_tree_543



type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var res int

func diameterOfBinaryTree(root *TreeNode) int {
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := dfs(root.Left)
	r := dfs(root.Right)
	path := max(l, r)
	res = max(l+r, res)
	return path + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}


