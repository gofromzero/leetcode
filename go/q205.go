package program

import "bytes"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	return reStr(s) == reStr(t)
}

func reStr(s string) string {
	var num = 'a'
	var bt bytes.Buffer
	m := make(map[int32]int32)
	for _, v := range s {
		val, ok := m[v]
		if !ok {
			num++
			m[v] = num
			val = num
		}
		bt.WriteRune(val)
	}
	return bt.String()
}
