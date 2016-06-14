package oj

import (
	"strings"
	"unicode"
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	n := len(s)
	if n == 0 {
		return false
	}
	if s[0] == '-' {
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	n = len(s)
	if n == 0 {
		return false
	}

	var isNumber = false
	i := 0
	for i < n && unicode.IsDigit(rune(s[i])) {
		i++
		isNumber = true
	}
	if i < n && s[i] == '.' {
		i++
		for i < n && unicode.IsDigit(rune(s[i])) {
			i++
			isNumber = true
		}
	}
	if isNumber && i < n && s[i] == 'e' {
		i++
		isNumber = false
		if i < n && (s[i] == '-' || s[i] == '+') {
			i++
		}
		for i < n && unicode.IsDigit(rune(s[i])) {
			i++
			isNumber = true
		}
	}
	return isNumber && i == n
}
