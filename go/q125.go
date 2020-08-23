package program

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	rep, _ := regexp.Compile("[^0-9a-z]")
	s = rep.ReplaceAllString(s, "")

	l := 0
	r := len(s) - 1
	for l < r {
		if s[r] != s[l] {
			return false
		}
		r--
		l++
	}
	return true
}
