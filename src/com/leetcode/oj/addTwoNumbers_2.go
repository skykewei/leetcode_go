package oj

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import ()

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{
		Val:  -1,
		Next: nil,
	}
	pret := ret
	counter := 0
	for l1 != nil && l2 != nil {
		tmp := l1.Val + l2.Val + counter
		pret.Next = &ListNode{
			Val:  tmp % 10,
			Next: nil,
		}
		counter = tmp / 10
		pret = pret.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		tmp := l1.Val + counter
		pret.Next = &ListNode{
			Val:  tmp % 10,
			Next: nil,
		}
		counter = tmp / 10
		pret = pret.Next
		l1 = l1.Next
	}
	for l2 != nil {
		tmp := l2.Val + counter
		pret.Next = &ListNode{
			Val:  tmp % 10,
			Next: nil,
		}
		counter = tmp / 10
		pret = pret.Next
		l2 = l2.Next
	}
	if counter > 0 {
		pret.Next = &ListNode{
			Val:  counter,
			Next: nil,
		}
	}
	return ret.Next
}
