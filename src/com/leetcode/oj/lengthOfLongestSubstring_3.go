package oj

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	s = strings.TrimSpace(s)
	var begin, i, n, mlen int
	begin = 0
	n = len(s)
	mlen = -1
	existed := make(map[byte]int)
	for i = 0; i < n; i++ {

		if val, ok := existed[s[i]]; ok {
			if val >= begin {
				if i-begin > mlen {
					mlen = i - begin
				}
				begin = val + 1
			}
		}
		existed[s[i]] = i
	}
	if i-begin > mlen {
		mlen = i - begin
	}
	return mlen
}
