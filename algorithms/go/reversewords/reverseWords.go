// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/1

package reversewords

import "strings"

func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	for i := 0; i < len(strs); i++ {
		strs[i] = reverseString([]byte(strs[i]))
	}

	return strings.Join(strs, " ")
}

func reverseString(s []byte) string {
	l, r := 0, len(s)-1
	for l <= r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	return string(s)
}
