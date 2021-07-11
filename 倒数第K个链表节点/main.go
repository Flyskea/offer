package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) findKthToTail(k int) *ListNode {
	if l == nil || k == 0 {
		return nil
	}
	p1 := l
	for i := 0; i < k-1; i++ {
		if p1.Next != nil {
			p1 = p1.Next
		} else {
			return nil
		}
	}

	p2 := l
	for p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}

func main() {
	l := ListNode{
		Next: nil,
		Val:  6,
	}
	fmt.Println(l.findKthToTail(1))
}
