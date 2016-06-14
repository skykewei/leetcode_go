package oj

import ()

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBal(left, right int) bool {
	if left == right {
		return true
	}
	if left-right == 1 {
		return true
	}
	if left-right == -1 {
		return true
	}
	return false
}
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := height(root.Left)
	right := height(root.Right)
	ret := left + 1
	if right > left {
		ret = right + 1
	}
	return ret
}
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := height(root.Left)
	right := height(root.Right)
	return isBal(left, right) && isBalanced(root.Left) && isBalanced(root.Right)
}
