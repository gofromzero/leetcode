package program

func mirrorReflection(p int, q int) int {
	m, n := p, q
	var r int
	for n > 0 {
		r = m % n
		m = n
		n = r
	}
	if (p/m)%2 == 0 {
		return 2
	} else if (q/m)%2 == 0 {
		return 0
	}

	return 1
}
