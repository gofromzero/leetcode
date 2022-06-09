// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/9

package maxareaofisland

var ( //表示当前位置的下，右，左，上四个方位
	dx = []int{1, 0, 0, -1}
	dy = []int{0, 1, -1, 0}
)

func maxAreaOfIsland(image [][]int) int {
	var max int
	n, m := len(image), len(image[0])
	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			if image[x][y] == 1 {
				temp := 1
				queue := [][]int{}
				queue = append(queue, []int{x, y})
				image[x][y] = 2
				for i := 0; i < len(queue); i++ { //对于队列中的所有点进行染色
					cell := queue[i]
					for j := 0; j < 4; j++ {
						mx := cell[0] + dx[j]
						my := cell[1] + dy[j]
						if mx >= 0 && mx < n && my >= 0 && my < m && image[mx][my] == 1 {
							image[mx][my] = 2
							queue = append(queue, []int{mx, my})
							temp++
						}
					}
				}
				if temp > max {
					max = temp
				}
			}
		}
	}

	return max
}
