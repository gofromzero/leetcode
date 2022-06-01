// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/5/26

package lengthoflongestsubstring

func lengthOfLongestSubstring(s string) int {
	var result int
	m := make(map[uint8]int)
	low := 0
	for i := 0; i < len(s); i++ {
		if index, ok := m[s[i]]; ok {
			low = max(index, low)
		}

		result = max(result, i-low+1)
		m[s[i]] = i + 1
	}

	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
