package decodeMessage

func decodeMessage(key string, message string) string {

	var m = map[int32]int32{}
	var k = 'a'
	for _, val := range key {
		if _, ok := m[val]; !ok && val != ' ' {
			m[val] = k
			k++
		}
	}

	result := []int32{}
	for _, val := range message {
		var temp int32 = ' '
		if val != ' ' {
			temp = m[val]
		}
		result = append(result, temp)
	}
	return string(result)
}
