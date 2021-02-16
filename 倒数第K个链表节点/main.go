package main

import (
	"fmt"
)

type list struct {
	next *list
	num  int
}

func (l *list) findKthToTail(k int) *list {
	if l == nil || k == 0 {
		return nil
	}
	p1 := l
	for i := 0; i < k-1; i++ {
		if p1.next != nil {
			p1 = p1.next
		} else {
			return nil
		}
	}

	p2 := l
	for p1.next != nil {
		p1 = p1.next
		p2 = p2.next
	}
	return p2
}

func main() {
	l := list{
		next: nil,
		num:  6,
	}
	fmt.Println(l.findKthToTail(1))
}
