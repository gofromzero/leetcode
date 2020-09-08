package program

func maximum69Number(num int) int {
	l := 1
	for ; l < num; l *= 10 {
	}
	for ; l > 0; l /= 10 {
		k := (num % l) * 10 / l
		if k == 6 {
			num += 3 * l / 10
			break
		}
	}
	return num
}
