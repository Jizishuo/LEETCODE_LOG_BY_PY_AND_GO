package Substructure_of_tree_26


type TreeNode struct {
    Val int
	Left *TreeNode
    Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A==nil && B==nil {
		return true
	}
	if A == nil || B==nil {
		return false
	}
	return same(A, B)|| isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
func same(a, b *TreeNode) bool {
	if b==nil{
		return true
	}
	if a==nil{
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return same(a.Left,b.Left) && same(a.Right, b.Right)
}



