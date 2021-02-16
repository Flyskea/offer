package main

import (
	"fmt"
)

type RandomListNode struct {
	Label  int
	Next   *RandomListNode
	Random *RandomListNode
}

func cloneNodes(head *RandomListNode) {
	var node *RandomListNode
	node = head
	for node != nil {
		cloned := &RandomListNode{
			Label: node.Label,
			Next:  node.Next,
		}
		node.Next = cloned
		node = cloned.Next
	}
}

func connectSiblingNodes(head *RandomListNode) {
	var node *RandomListNode
	node = head
	for node != nil {
		cloned := node.Next
		if node.Random != nil {
			cloned.Random = node.Random.Next
		}
		node = cloned.Next
	}
}

func reconnectNodes(head *RandomListNode) *RandomListNode {
	var node *RandomListNode
	node = head
	var clonedHead *RandomListNode
	var clonedNode *RandomListNode
	if node != nil {
		clonedNode = node.Next
		clonedHead = clonedNode
		node.Next = clonedNode.Next
		node = node.Next
	}
	for node != nil {
		clonedNode.Next = node.Next
		clonedNode = clonedNode.Next
		node.Next = clonedNode.Next
		node = node.Next
	}
	return clonedHead
}

func Clone(head *RandomListNode) *RandomListNode {
	cloneNodes(head)
	connectSiblingNodes(head)
	return reconnectNodes(head)
}

func main() {
	l := RandomListNode{Label: 2}
	l1 := &RandomListNode{Label: 3}
	l.Next = l1
	l2 := &RandomListNode{Label: 23}
	l.Random = l2
	l1.Next = l2
	fmt.Println(Clone(&l))
}
