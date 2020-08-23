package program

import "strings"

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
