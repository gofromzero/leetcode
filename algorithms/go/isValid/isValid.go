package isValid

var strs = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	arr := make([]rune, 0)

	var temp rune
	for _, val := range s {
		if len(arr) != 0 {
			temp = strs[val]
			if temp == arr[len(arr)-1] {
				arr = arr[0 : len(arr)-1]
				continue
			}
		}
		arr = append(arr, val)
	}

	return len(arr) == 0
}
