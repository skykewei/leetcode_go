package oj

import "strings"

func strStr2(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func strStr(haystack, needle string) int {
	if needle == "" {
		return 0
	}

	var hl, nl int = len(haystack), len(needle)
	if hl < nl {
		return -1
	}
	max := hl - nl

	ns := needle[0:]
	for i := 0; i <= max; i++ {
		hs := haystack[i : i+nl]
		if hs == ns {
			return i
		}
	}
	return -1
}
