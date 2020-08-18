package program

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	l := 2*numRows - 2

	arr := make([]string, numRows)
	for k, v := range s {
		temp := k % l
		if temp < numRows {
			arr[temp] += string(v)
		} else {
			arr[l-temp] += string(v)
		}
	}

	return strings.Join(arr, "")
}
