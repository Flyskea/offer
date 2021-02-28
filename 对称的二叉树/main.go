package main

// TreeNode ..
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetricalHelper(p1, p2 *TreeNode) bool {
	if p1 == nil {
		return p2 == nil
	}
	if p2 == nil {
		return false
	}
	if p1.Val != p2.Val {
		return false
	}
	return isSymmetricalHelper(p1.Left, p2.Right) && isSymmetricalHelper(p1.Right, p2.Left)
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param pRoot TreeNode类
 * @return bool布尔型
 */
func isSymmetrical(pRoot *TreeNode) bool {
	return isSymmetricalHelper(pRoot, pRoot)
}
