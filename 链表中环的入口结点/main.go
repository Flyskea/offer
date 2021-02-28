package main

// ListNode ..
type ListNode struct {
	Val  int
	Next *ListNode
}

func meetingNode(pHead *ListNode) *ListNode {
	if pHead == nil {
		return nil
	}
	slow := pHead.Next
	if slow == nil {
		return nil
	}
	fast := slow.Next
	for fast != nil && slow != nil {
		if fast == slow {
			return fast
		}
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	return nil
}

// EntryNodeOfLoop ..
func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	// 找到环里节点
	mNode := meetingNode(pHead)
	if mNode == nil {
		return nil
	}

	nodesInLoop := 1
	p := mNode
	// 计算环节点数
	for p.Next != mNode {
		p = p.Next
		nodesInLoop++
	}
	p = pHead
	for i := 0; i < nodesInLoop; i++ {
		p = p.Next
	}
	p2 := pHead
	for p != p2 {
		p = p.Next
		p2 = p2.Next
	}
	return p
}
