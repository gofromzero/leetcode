package program

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	l := len(startTime)
	var count int
	for i := 0; i < l; i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			count++
		}
	}
	return count
}
