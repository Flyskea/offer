package main

import (
	"fmt"
)

func isEven(num int) bool {
	return num&1 == 0
}

func recoderOddEven(array []int, f func(int) bool) {
	l := len(array)
	if l == 0 {
		return
	}
	begin, end := 0, l-1
	for begin < end {
		for begin < end && !f(array[begin]) {
			begin++
		}
		for begin < end && f(array[end]) {
			end--
		}
		if begin < end {
			tmp := array[begin]
			array[begin] = array[end]
			array[end] = tmp
		}
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 7, 9, 6, 2}
	recoderOddEven(a, isEven)
	fmt.Println(a)
}
