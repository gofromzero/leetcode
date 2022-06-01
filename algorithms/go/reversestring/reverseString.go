// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/1

package reversestring

func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l <= r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
