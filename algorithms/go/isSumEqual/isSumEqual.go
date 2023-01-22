package isSumEqual

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return count(firstWord)+count(secondWord) == count(targetWord)
}

func count(str string) int32 {
	var result int32
	for _, val := range str {
		result *= 10
		result += val - 'a'
	}

	return result
}
