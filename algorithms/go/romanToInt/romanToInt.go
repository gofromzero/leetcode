package romanToInt

var roman = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	strs := []rune(s)
	pre := 0
	var result int
	for i := len(strs) - 1; i >= 0; i-- {
		temp := roman[strs[i]]
		if temp >= pre {
			result += temp
		} else {
			result -= temp
		}
		pre = temp
	}
	return result
}
