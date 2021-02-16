package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param pre int整型一维数组
 * @param vin int整型一维数组
 * @return TreeNode类
 */
func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	for k := range vin {
		if vin[k] == pre[0] { //中序遍历 root (index=k)
			return &TreeNode{
				Val: pre[0],
				//Val: inorder[k],
				Left:  reConstructBinaryTree(pre[1:k+1], vin[0:k]),
				Right: reConstructBinaryTree(pre[k+1:], vin[k+1:]),
			}
		}
	}
	return nil
}
