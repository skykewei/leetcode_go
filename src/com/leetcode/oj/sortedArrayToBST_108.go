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
func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil {
		return nil
	}
	n := len(nums)
	if n == 0 {
		return nil
	}
	mid := n / 2
	ret := &TreeNode{
		Val:   nums[mid],
		Left:  nil,
		Right: nil,
	}
	ret.Left = sortedArrayToBST(nums[0:mid])
	ret.Right = sortedArrayToBST(nums[mid+1:])
	return ret
}
