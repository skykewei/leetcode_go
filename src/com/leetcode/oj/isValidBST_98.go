package oj

import ()

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	pLeft := root.Left
	for pLeft != nil && pLeft.Right != nil {
		pLeft = pLeft.Right
	}
	if pLeft != nil && !(pLeft.Val < root.Val) {
		return false
	}
	pRight := root.Right
	for pRight != nil && pRight.Left != nil {
		pRight = pRight.Left
	}
	if pRight != nil && !(pRight.Val > root.Val) {
		return false
	}
	return isValidBST(root.Left) && isValidBST(root.Right)
}
