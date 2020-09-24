package program

func restoreString(s string, indices []int) string {
	bt := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		bt[indices[i]] = s[i]
	}
	return string(bt)
}
