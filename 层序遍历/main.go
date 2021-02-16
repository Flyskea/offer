package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrintFromTopToBottom(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		res = append(res, v.Val)
		if v.Left != nil {
			q = append(q, v.Left)
		}
		if v.Right != nil {
			q = append(q, v.Right)
		}
	}
	return res
}

func main() {
	root := &TreeNode{Val: 1}
	left := &TreeNode{Val: 2}
	root.Left = left
	fmt.Println(PrintFromTopToBottom(root))
}
