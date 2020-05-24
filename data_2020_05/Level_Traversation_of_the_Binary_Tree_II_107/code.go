package Level_Traversation_of_the_Binary_Tree_II_107


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	for root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) >0 {
		l := len(queue)
		list := make([]int, 0)
		for i :=0;i<l;i++ {
			node := queue[i]
			list = append(list, node.Val)
			if node.Left !=nil {
				queue =append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append([][]int{list}, res...)
		queue = queue[l:]
	}
	return res
}

