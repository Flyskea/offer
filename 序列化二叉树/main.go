package main

// TreeNode ..
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Serialization1 ..
func Serialization1(root *TreeNode, arr *[]int) []int {
	if root == nil {
		*arr = append(*arr, -1)
		return *arr
	}
	*arr = append(*arr, root.Val)
	Serialization1(root.Left, arr)
	Serialization1(root.Right, arr)
	return *arr

}

// DeSerialization ..
func DeSerialization(arr *[]int) *TreeNode {
	cur := (*arr)[0]
	*arr = (*arr)[1:]
	if cur == -1 {
		return nil
	}
	root := &TreeNode{Val: cur}
	root.Left = DeSerialization(arr)
	root.Right = DeSerialization(arr)
	return root
}

// Serialize ..
func Serialize(root *TreeNode) *TreeNode {
	// 前序遍历进行序列化
	arr := make([]int, 0)
	a := Serialization1(root, &arr)
	root = DeSerialization(&a)
	return root
}
