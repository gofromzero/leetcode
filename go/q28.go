package program

import (
	"strings"
)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func strStrInternal(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}

	origintIndex := 0
	aimIndex := 0
	for aimIndex < len(needle) {
		if origintIndex > len(haystack)-1 {
			return -1
		}
		if haystack[origintIndex] == needle[aimIndex] {
			origintIndex++
			aimIndex++
		} else {
			nextCharIndex := origintIndex - aimIndex + len(needle)
			if nextCharIndex < len(haystack) {
				step := strings.LastIndex(needle, string(haystack[nextCharIndex]))
				if step == -1 {
					origintIndex = nextCharIndex + 1
				} else {
					origintIndex = nextCharIndex - step
				}
				aimIndex = 0
			} else {
				return -1
			}
		}
	}

	return origintIndex - aimIndex
}

func strStrKMP(haystack string, needle string) int {
	next := gnext(needle)
	return kmp(haystack, needle, next)
}

func gnext(needle string) []int {
	j := -1
	i := 0
	next := make([]int, len(needle)+1)
	next[i] = j
	for i < len(needle) {
		for j >= 0 && needle[i] != needle[j] {
			j = next[j]
		}
		i++
		j++
		next[i] = j
	}
	return next
}

func kmp(haystack string, needle string, next []int) int {
	l := len(haystack)
	l2 := len(needle)
	var i, j int
	for i < l && j < l2 {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == l2 {
		return i - j
	}

	return -1
}
