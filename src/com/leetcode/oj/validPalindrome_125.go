package oj

import (
	"unicode"
)

//oj https://leetcode.com/problems/valid-palindrome/
func isPalindrome_125(s string) bool {
	r := s
	i, j := 0, len(r)-1
	for i < j {
		for i < j && !(unicode.IsDigit(rune(r[i])) || unicode.IsLetter(rune(r[i]))) {
			i++
		}
		for i < j && !(unicode.IsDigit(rune(r[j])) || unicode.IsLetter(rune(r[j]))) {
			j--
		}
		if unicode.ToLower(rune(r[i])) != unicode.ToLower(rune(r[j])) {
			return false
		}
		i++
		j--
	}
	return true
}
