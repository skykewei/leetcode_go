package oj

import ()

// 121 https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	curMin := prices[0]
	profit := 0
	for i := 1; i < n; i++ {
		profit = maxInt(profit, prices[i]-curMin)
		curMin = minInt(curMin, prices[i])
	}
	return profit
}
func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
