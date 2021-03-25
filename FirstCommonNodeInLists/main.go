package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil || pHead2 == nil {
		return nil
	}
	nLength1 := listLen(pHead1)
	nLength2 := listLen(pHead2)
	nDifLen := nLength1 - nLength2
	pLong := pHead1
	pShort := pHead2
	if nLength1 < nLength2 {
		pLong = pHead2
		pShort = pHead1
		nDifLen = nLength2 - nLength1
	}
	for i := 0; i < nDifLen; i++ {
		pLong = pLong.Next
	}
	for pLong != nil && pShort != nil && pLong != pShort {
		pLong = pLong.Next
		pShort = pShort.Next
	}
	return pLong
}

func listLen(head *ListNode) int {
	var length int
	p := head
	for p != nil {
		length++
		p = p.Next
	}
	return length
}
