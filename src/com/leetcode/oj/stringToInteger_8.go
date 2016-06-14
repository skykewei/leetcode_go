package oj

import (
	"strings"
	"unicode"
)

func myAtoi(str string) int {

	const (
		INT_MAX = 2147483647
		INT_MIN = -2147483648
	)
	str2 := strings.TrimSpace(str)
	n := len(str2)
	if n == 0 {
		return 0
	}
	i := 0
	var signed int64 = 1
	if i < n && str2[i] == '-' {
		signed = -1
		i++
	} else if i < n && str2[i] == '+' {
		signed = 1
		i++
	}
	if i < n && !unicode.IsDigit(rune(str2[i])) {
		return 0
	}

	var ret int64 = 0
	var base int64 = 10
	for i < n && unicode.IsDigit(rune(str2[i])) {
		ret = ret*base + int64(str2[i]) - int64('0')
		if signed == 1 && ret > INT_MAX {
			return INT_MAX
		} else if signed == -1 && ret > (-INT_MIN) {
			return INT_MIN
		}
		i++
	}
	return int(ret) * int(signed)

}
