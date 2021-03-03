package main

// TreeNode ..
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// KthNode ..
func KthNode(pRoot *TreeNode, k int) *TreeNode {
	if pRoot == nil || k == 0 {
		return nil
	}
	return kthNodeCore(pRoot, &k)
}

func kthNodeCore(pRoot *TreeNode, k *int) *TreeNode {
	var target *TreeNode
	if pRoot.Left != nil {
		target = kthNodeCore(pRoot.Left, k)
	}
	if target == nil {
		if *k == 1 {
			target = pRoot
		}
		*k--
	}
	if target == nil && pRoot.Right != nil {
		target = kthNodeCore(pRoot.Right, k)
	}
	return target
}
