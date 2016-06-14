package oj

import ()

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, -1
	len := len(s)
	for i := 0; i < len; i++ {
		n1 := expandPal(s, i, i)
		n2 := expandPal(s, i, i+1)
		nm := n1
		if n2 > nm {
			nm = n2
		}
		if nm > end-start {
			start = i - (nm-1)/2
			end = i + nm/2
		}
	}
	return s[start : end+1]
}

func expandPal(s string, start int, end int) int {
	i, j := start, end
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return j - i - 1
}
