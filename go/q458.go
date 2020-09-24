package program

import "math"

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	state := minutesToTest/minutesToDie + 1
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(state))))
}
