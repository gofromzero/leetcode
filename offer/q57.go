package offer

func findContinuousSequence(target int) [][]int {
	result := make([][]int, 0)
	sum := 0
	left, right := 1, 1
	for right < target {
		if sum > target {
			sum -= left
			left++
		} else if sum < target {
			sum += right
			right++
		} else {
			result = append(result, makeArr(left, right-1))
			sum -= left
			left++
		}
	}
	return result
}

func makeArr(start int, end int) []int {
	l := end - start + 1
	res := make([]int, l)
	for i := range res {
		res[i] = start
		start++
	}
	return res
}
