package longestPalindrome

func longestPalindrome(s string) int {
	var count = map[rune]int{}
	for _, val := range s {
		count[val]++
	}
	l := 0 //双
	r := 0 // 单
	num := 0
	for _, val := range count {
		if val%2 == 0 {
			l += val
		} else {
			r += val
			num++
		}
	}
	if num != 0 {
		num--
	}

	return l + r - num
}
