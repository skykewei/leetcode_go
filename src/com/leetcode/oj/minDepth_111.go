package oj

import (
	"container/list"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// using Recursion
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	}

	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)

	ret := left + 1
	if right < left {
		ret = right + 1
	}
	return ret
}

// using breadth-first traversal
func minDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := list.New()
	depth := 1
	rightmost := root
	if rightmost == nil {
		rightmost = root.Left
	}
	q.PushBack(root)
	for q.Len() != 0 {
		var el *list.Element = q.Front()
		var t *TreeNode = (el.Value).(*TreeNode)
		if t.Left == nil && t.Right == nil {
			break
		}
		if t.Left != nil {
			q.PushBack(t.Left)
		}
		if t.Right != nil {
			q.PushBack(t.Right)
		}
		if t == rightmost {
			depth++
			rightmost = t.Right
			if rightmost == nil {
				rightmost = t.Left
			}
		}
		q.Remove(el)
	}
	return depth
}

// using slice
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var q []*TreeNode
	depth := 1
	rightmost := root
	if rightmost == nil {
		rightmost = root.Left
	}
	q = append(q, root)
	for len(q) != 0 {
		t := q[0]
		q = q[1:]
		if t.Left == nil && t.Right == nil {
			break
		}
		if t.Left != nil {
			q = append(q, t.Left)
		}
		if t.Right != nil {
			q = append(q, t.Right)
		}
		if t == rightmost {
			depth++
			rightmost = t.Right
			if rightmost == nil {
				rightmost = t.Left
			}
		}

	}
	return depth
}
