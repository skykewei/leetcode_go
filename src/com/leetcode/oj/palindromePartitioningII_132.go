package oj

import ()

// 132. Palindrome Partitioning II
// https://leetcode.com/problems/palindrome-partitioning-ii/
func minCut(s string) int {
	n := len(s)
	if n == 0 {
		return -1
	}
	mc := make([]int, n+1)
	pt := make([]bool, n*n)
	pp := make([][]bool, n)
	for i := 0; i < n; i++ {
		pp[i] = pt[i*n : (i+1)*n]
		mc[i] = n - i - 1
	}
	mc[n] = -1

	for i := n - 1; i >= 0; i-- {
		for j := i; j <= n-1; j++ {
			if s[i] == s[j] && (j-i < 2 || pp[i+1][j-1] == true) {
				pp[i][j] = true
			}
			if pp[i][j] == true && mc[i] > mc[j+1]+1 {
				mc[i] = mc[j+1] + 1
			}
		}
	}
	return mc[0]
}
