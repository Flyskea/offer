package main

// TreeLinkNode next is parent node
type TreeLinkNode struct {
	Val   int
	Left  *TreeLinkNode
	Right *TreeLinkNode
	Next  *TreeLinkNode
}

// GetNext ..
func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
	if pNode == nil {
		return nil
	}
	var next *TreeLinkNode
	if pNode.Right != nil {
		right := pNode.Right
		for right.Left != nil {
			right = right.Left
		}
		next = right
	} else if pNode.Next != nil {
		current := pNode
		parent := pNode.Next
		for parent != nil && current == parent.Right {
			current = parent
			parent = parent.Next
		}
		next = parent
	}
	return next
}
