// Package numberOfArrays
// @Author Spume
// @Date 2025-04-21 10:09

package numberOfArrays

func numberOfArrays(differences []int, lower int, upper int) int {
	pre := lower

	l := lower
	r := lower
	for _, val := range differences {
		pre = pre + val
		if l > pre {
			l = pre
		}
		if r < pre {
			r = pre
		}
	}

	width := upper - lower
	if width < r-l {
		return 0
	} else {
		return width - (r - l)
	}

}
