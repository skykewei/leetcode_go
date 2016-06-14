package oj

import ()

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	if len(lists) == 0 {
		return nil
	}
	begin := 0
	end := len(lists) - 1
	if end == 0 {
		return lists[0]
	}
	if begin+1 == end {
		return merge2Lists(lists[begin], lists[end])
	}

	mid := (begin + end) / 2
	first := mergeKLists(lists[begin:mid])
	second := mergeKLists(lists[mid:])
	return merge2Lists(first, second)

}
func merge2Lists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var retH *ListNode = nil
	var retT *ListNode = nil
	p1 := l1
	p2 := l2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			if retT == nil {
				retH = p1
				retT = p1
			} else {
				retT.Next = p1
				retT = retT.Next
			}
			p1 = p1.Next
		} else {
			if retT == nil {
				retH = p2
				retT = p2
			} else {
				retT.Next = p2
				retT = retT.Next
			}
			p2 = p2.Next
		}
	}
	if p1 != nil {
		retT.Next = p1
	}
	if p2 != nil {
		retT.Next = p2
	}
	return retH
}
