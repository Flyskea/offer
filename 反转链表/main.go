package main

import (
	"fmt"
)

type list struct {
	next *list
	num  int
}

func (l *list) reverseListRecursive() *list {
	if l == nil || l.next == nil {
		return l
	}
	p := l.next.reverseListRecursive()
	l.next.next = l
	l.next = nil
	return p
}

func (l *list) reverseList() {
	if l == nil {
		return
	}
	p := l
	var pre *list
	for p != nil {
		next := p.next
		p.next = pre
		pre = p
		p = next
	}
	l = p
}

func main() {
	l := &list{
		next: nil,
		num:  6,
	}
	k := list{
		next: l,
		num:  8,
	}
	k.reverseListRecursive()
	fmt.Println(l)
	fmt.Printf("%p\n", &k)
	var res *list
	fmt.Println(res)
}
