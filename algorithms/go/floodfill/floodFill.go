// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/9

package floodfill

var ( //表示当前位置的下，右，左，上四个方位
	dx = []int{1, 0, 0, -1}
	dy = []int{0, 1, -1, 0}
)

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	curColor := image[sr][sc]
	if curColor == newColor {
		return image
	}
	n, m := len(image), len(image[0])
	queue := [][]int{}
	queue = append(queue, []int{sr, sc})
	image[sr][sc] = newColor
	for i := 0; i < len(queue); i++ { //对于队列中的所有点进行染色
		cell := queue[i]
		for j := 0; j < 4; j++ {
			mx := cell[0] + dx[j]
			my := cell[1] + dy[j]
			if mx >= 0 && mx < n && my >= 0 && my < m && image[mx][my] == curColor {
				image[mx][my] = newColor
				queue = append(queue, []int{mx, my})
			}
		}
	}
	return image
}
