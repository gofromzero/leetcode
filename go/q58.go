package program

func lengthOfLastWord(s string) int {
	l := len(s)

	isOver := false
	num := 0
	for i := l - 1; i >= 0; i-- {
		if isChar(rune(s[i])) {
			num++
			if !isOver {
				isOver = true
			}
		} else if isOver {
			break
		}
	}

	return num
}

func isChar(a rune) bool {
	return (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z')
}
