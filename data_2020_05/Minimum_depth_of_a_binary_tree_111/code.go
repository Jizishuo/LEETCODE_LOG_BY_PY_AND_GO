package Minimum_depth_of_a_binary_tree_111


// Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var num int

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	num = 1<<63-1
	mmm(root, 1)
	return num


	
}

func mmm (root *TreeNode, n int) {
	if root.Left == nil && root.Right == nil {
		if n < num {
			num = n
		}
		return
	}
	if root.Left != nil {
		mmm(root.Left, n+1)
	}
	if root.Right != nil {
		mmm(root.Right, n+1)
	}

}