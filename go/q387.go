package program

func firstUniqChar(s string) int {
	m := make(map[rune]int, len(s))
	for _, v := range s {
		m[v]++
	}
	for i, v := range s {
		if value, ok := m[v]; ok {
			if value == 1 {
				return i
			}
		}
	}

	return -1
}

func firstUniqCharBetter(s string) int {
	var arr [26]int
	for i, k := range s {
		arr[k-'a'] = i
	}
	for i, k := range s {
		if i == arr[k-'a'] {
			return i
		}
		arr[k-'a'] = -1
	}
	return -1
}
