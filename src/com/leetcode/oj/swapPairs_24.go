package oj

import ()

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	ret := &ListNode{
		-1, nil,
	}
	pret := ret
	p1 := head
	if p1.Next == nil {
		return head
	}
	p2 := p1.Next

	for p2 != nil && p1 != nil {
		tmp := p2.Next
		pret.Next = p2
		p2.Next = p1
		p1.Next = nil
		pret = p1
		p1 = tmp
		if p1 != nil {
			p2 = p1.Next
			if p2 == nil {
				pret.Next = p1
			}
		}
	}

	return ret.Next
}
