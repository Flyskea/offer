package main

// ListNode ..
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
func deleteDuplication(pHead *ListNode) *ListNode {
	if pHead == nil {
		return nil
	}
	var pre *ListNode
	p := pHead
	for p != nil {
		next := p.Next
		needDelete := false
		if next != nil && next.Val == p.Val {
			needDelete = true
		}
		if !needDelete {
			pre = p
			p = p.Next
		} else {
			value := p.Val
			pToBeDel := p
			for pToBeDel != nil && pToBeDel.Val == value {
				next = pToBeDel.Next
				pToBeDel = nil
				pToBeDel = next
			}
			if pre == nil {
				pHead = next
			} else {
				pre.Next = next
			}
			p = next
		}
	}
	return pHead
}
