package isIsomorphic

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var s2t = map[byte]byte{}
	var t2s = map[byte]byte{}

	for i := 0; i < len(s); i++ {
		x := s[i]
		y := t[i]

		if val, ok := s2t[x]; ok && val != y {
			return false
		}

		if val, ok := t2s[y]; ok && val != x {
			return false
		}

		s2t[x] = y
		t2s[y] = x
	}

	return true
}
