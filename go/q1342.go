package program

func numberOfSteps(num int) int {
	count := 0
	for num != 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
		count++
	}
	return count
}
