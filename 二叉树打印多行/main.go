package main

// TreeNode ..
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Print ..
func Print(pRoot *TreeNode) [][]int {
	if pRoot == nil {
		return nil
	}
	var res [][]int
	queue := []*TreeNode{pRoot}
	for len(queue) > 0 {
		vals := []int{}
		for n := len(queue); n > 0; n++ {
			val := queue[0]
			queue = queue[1:]
			vals = append(vals, val.Val)
			if val.Left != nil {
				queue = append(queue, val.Left)
			}
			if val.Right != nil {
				queue = append(queue, val.Right)
			}
		}
		res = append(res, vals)
	}
	return res
}
