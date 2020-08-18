package program

type pos struct {
	sr int
	sc int
}

var dx = []int{0, 0, -1, 1}
var dy = []int{1, -1, 0, 0}
var count int

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	if newColor == oldColor {
		return image
	}
	dp := []pos{{sr, sc}}
	image[sr][sc] = newColor
	for len(dp) != 0 {
		count++
		sr := dp[0].sr
		sc := dp[0].sc
		for i := 0; i < 4; i++ {
			x, y := sr+dx[i], sc+dy[i]
			if x >= 0 && x < len(image) && y >= 0 && y < len(image[0]) && image[x][y] == oldColor {
				image[x][y] = newColor
				dp = append(dp, pos{x, y})
			}

		}
		dp = dp[1:]
	}
	return image
}
