package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// FindPath ...
func FindPath(root *TreeNode, expectNumber int) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	Find(root, expectNumber, 0, []int{}, &res)
	return res
}

// Find ...
func Find(root *TreeNode, expectNumber, current int, a []int, path *[][]int) {
	if root == nil {
		return
	}
	a = append(a, root.Val)
	current += root.Val
	if root.Left == nil && root.Right == nil && current == expectNumber {
		*path = append(*path, a)
	}
	if root.Left != nil {
		Find(root.Left, expectNumber, current, a, path)
	}
	if root.Right != nil {
		Find(root.Right, expectNumber, current, a, path)
	}
	a = a[:len(a)-1]
}
