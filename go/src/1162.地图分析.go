/*
 * @lc app=leetcode.cn id=1162 lang=golang
 *
 * [1162] 地图分析
 */

// @lc code=start
func maxDistance(grid [][]int) int {
	// DP 分两阶段进行 初始化 所有陆地值都为0，海洋值为maxint
	// 接下来的计算只针对海洋值
	// 第一阶段，计算从上方和左方到达的距离
	//		f[i][j] = min(f[i-1][j], minf[i][j-1]) + 1
	// 第二阶段，计算从右方和下方到达的距离
	//    f[i][j] = min(f[i+1][j], minf[i][j+1]) + 1
	n := len(grid)
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				f[i][j] = 0
			} else {
				f[i][j] = n * 2
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if f[i][j] > 0 {
				if i - 1 >= 0 && f[i-1][j] + 1 < f[i][j] {
					f[i][j] = f[i-1][j] + 1
				}
				if j - 1 >= 0 && f[i][j-1] + 1 < f[i][j] {
					f[i][j] = f[i][j-1] + 1
				}
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if f[i][j] > 0 {
				if i + 1 < n && f[i+1][j] + 1 < f[i][j] {
					f[i][j] = f[i+1][j] + 1
				}
				if j + 1 < n && f[i][j+1] + 1 < f[i][j] {
					f[i][j] = f[i][j+1] + 1
				}
			}
		}
	}

	max := -1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 && f[i][j] > max {
				max = f[i][j]
			}
		}
	}

	if max == 2 * n {
		max = -1
	}
	return max
}

func maxDistance2(grid [][]int) int {
	n := len(grid)
	queue := []int{}
	counter := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, i*n + j)
				counter[i*n+j] = 0
			}
		}
	}

	if len(queue) == n * n || len(queue) == 0 {
		return -1
	}

	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		x, y := cur / n, cur % n
		count := counter[x*n+y]
		for i := 0; i < 4; i++ {
			nx, ny := x + dx[i], y + dy[i]
			if nx < 0 || nx == n || ny < 0 || ny == n || grid[nx][ny] == 1 {
				continue
			}

			if _, ok := counter[nx*n+ny]; ok {
				continue
			} 
			counter[nx*n+ny] = count + 1
			queue = append(queue, nx*n+ny)
		}
	}

	max := 0
	for _, v := range counter {
		if v > max {
			max = v
		}
	}
	return max
}
// @lc code=end

