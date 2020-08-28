package program

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
