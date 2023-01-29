package countAsterisks

func countAsterisks(s string) int {
	var result int
	var lnum int
	for _, val := range s {
		if val == '|' {
			lnum++
		} else if val == '*' && lnum%2 == 0 {
			result++
		}
	}
	return result
}
