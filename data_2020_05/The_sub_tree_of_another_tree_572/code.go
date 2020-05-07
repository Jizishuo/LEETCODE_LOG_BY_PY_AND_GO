
package The_sub_tree_of_another_tree_572

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {

	if s==nil {
		return false
	}
	return issame(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}
func issame(s *TreeNode, t *TreeNode) bool {
	if s ==nil && t== nil {
		return true
	}
	if s==nil ||t ==nil {
		return false
	}
	if s.Val==t.Val {
		return issame(s.Left, t.Left) && issame(s.Right, t.Right)
	}
	return false
}