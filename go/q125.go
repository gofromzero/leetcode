package program

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l := 0
	r := len(s) - 1
	for l < r {
		if !check(rune(s[l])) {
			l++
			continue
		}
		if !check(rune(s[r])) {
			r--
			continue
		}

		if s[r] != s[l] {
			return false
		}
		r--
		l++
	}
	return true
}

func check(v rune) bool {
	return ('0' <= v && v <= '9') || (v >= 'a' && v <= 'z')
}
