package main

import "fmt"

func multiply(A []int) []int {
	aLen := len(A)
	B := make([]int, aLen)
	tmp := 1
	for i := 0; i < aLen; i++ {
		B[i] = tmp
		tmp *= A[i]
	}
	tmp = 1
	for i := aLen - 1; i >= 0; i-- {
		B[i] *= tmp
		tmp *= A[i]
	}
	return B
}

func main() {
	fmt.Println(multiply([]int{1, 2, 3, 4, 5}))
}
