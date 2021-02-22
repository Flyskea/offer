package main

import (
	"container/list"
	"fmt"
)

// LastRemaining_Solution ..
func LastRemaining_Solution(n int, m int) int {
	if n < 1 || m < 1 {
		return -1
	}
	lists := list.New()
	for i := 0; i < n; i++ {
		lists.PushBack(i)
	}
	current := lists.Front()
	for lists.Len() > 1 {
		for i := 1; i < m; i++ {
			if current == lists.Back() {
				current = lists.Front()
				i++
				continue
			}
			current = current.Next()
			if current == lists.Back() {
				current = lists.Front()
				i++
			}
		}
		next := current.Next()
		lists.Remove(current)
		current = next
	}
	return current.Value.(int)
}

func main() {
	fmt.Println(LastRemaining_Solution(5, 2))
}
