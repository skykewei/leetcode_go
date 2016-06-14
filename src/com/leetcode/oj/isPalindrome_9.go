package oj

import ()

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	div := 1
	for x/div >= 10 {
		div *= 10
	}
	for x != 0 {
		l := x / div
		r := x % 10
		if l != r {
			return false
		}
		x %= div
		x /= 10
		div /= 100
	}
	return true
}
