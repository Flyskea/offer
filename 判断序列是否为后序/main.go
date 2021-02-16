package main

import (
	"fmt"
	"math"
)

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 2 {
		return true
	}
	head := len(postorder) - 1
	for head != 0 {
		popinter := 0
		for postorder[popinter] < postorder[head] {
			popinter++
		}
		for postorder[popinter] > postorder[head] {
			popinter++
		}
		if popinter != head {
			return false
		}
		head--
	}
	return true
}

func verifySquenceOfBSTStack(sequence []int) bool {
	stack, root := []int{}, math.MaxInt32
	for i := len(sequence) - 1; i > 0; i-- {
		if sequence[i] > root {
			return false
		}
		for len(stack) != 0 && sequence[i] < stack[len(stack)-1] {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, sequence[i])
	}
	return true
}

func verifySquenceOfBST(sequence []int) bool {
	if sequence == nil || len(sequence) == 0 {
		return false
	}
	length := len(sequence)
	root := sequence[length-1]
	i := 0
	for ; i < length-1; i++ {
		if sequence[i] > root {
			break
		}
	}
	j := i
	for ; j < length-1; j++ {
		if sequence[j] < root {
			return false
		}
	}
	left := true
	if i > 0 {
		left = verifySquenceOfBST(sequence[:i])
	}
	right := true
	if i < length-1 {
		right = verifySquenceOfBST(sequence[i : length-1])
	}
	return (left && right)
}

func main() {
	fmt.Println(verifySquenceOfBST([]int{5, 7, 26, 11, 10, 8}))
	fmt.Println(verifySquenceOfBSTStack([]int{5, 7, 26, 1, 11, 10, 8}))
	fmt.Println(verifyPostorder([]int{5, 7, 26, 1, 11, 10, 8}))
}
