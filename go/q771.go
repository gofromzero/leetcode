package program

func numJewelsInStones(J string, S string) int {
	jew := map[rune]bool{}
	for _, v := range J {
		jew[v] = true
	}
	var count int
	for _, v := range S {
		if jew[v] {
			count++
		}
	}
	return count
}
