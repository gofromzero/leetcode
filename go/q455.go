package program

import "sort"

func findContentChildren(g []int, s []int) int {
	if len(g) == 0 || len(s) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	var gi, sj, count int
	for gi < len(g) && sj < len(s) {
		if s[sj] >= g[gi] {
			count++
			gi++
		}
		sj++
	}
	return count
}
