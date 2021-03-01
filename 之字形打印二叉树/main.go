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
	level := make([][]*TreeNode, 2)
	var res [][]int
	current := 0
	next := 1
	level[0] = append(level[0], pRoot)
	for len(level[0]) != 0 || len(level[1]) != 0 {
		vals := []int{}
		for len(level[0]) != 0 || len(level[1]) != 0 {
			length := len(level[current]) - 1
			p := level[current][length]
			level[current] = level[current][:length]
			vals = append(vals, p.Val)
			if current == 0 {
				if p.Left != nil {
					level[next] = append(level[next], p.Left)
				}
				if p.Right != nil {
					level[next] = append(level[next], p.Right)
				}
			} else {
				if p.Right != nil {
					level[next] = append(level[next], p.Right)
				}
				if p.Left != nil {
					level[next] = append(level[next], p.Left)
				}
			}
			if len(level[current]) == 0 {
				current = 1 - current
				next = 1 - next
				break
			}
		}
		res = append(res, vals)
	}
	return res
}
