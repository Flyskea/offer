package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param pRoot1 TreeNode类
 * @param pRoot2 TreeNode类
 * @return bool布尔型
 */
func HasSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	if pRoot2 == nil || pRoot1 == nil {
		return false
	}

	return isSubtree(pRoot1, pRoot2) || HasSubtree(pRoot1.Left, pRoot2) || HasSubtree(pRoot1.Right, pRoot2)
}

func isSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	if pRoot2 == nil {
		return true
	}
	if pRoot1 == nil {
		return false
	}

	return pRoot1.Val == pRoot2.Val && isSubtree(pRoot1.Left, pRoot2.Left) && isSubtree(pRoot1.Right, pRoot2.Right)
}
