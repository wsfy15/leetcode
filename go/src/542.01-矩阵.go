/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */

// @lc code=start
func updateMatrix(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				res[i][j] = math.MaxInt32 >> 1
			}
		}
	}

	// 先判断左上方向来的
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if res[i][j] != 0 {
				if i >= 1 {
					res[i][j] = min(res[i][j], res[i-1][j] + 1)
				}
				if j >= 1 {
					res[i][j] = min(res[i][j], res[i][j-1] + 1)
				}
			}
		}
	}

		// 再判断右下方向来的
		for i := m - 1; i >= 0; i-- {
			for j := n - 1; j >= 0; j-- {
				if res[i][j] != 0 {
					if i <= m - 2 {
						res[i][j] = min(res[i][j], res[i+1][j] + 1)
					}
					if j <= n - 2 {
						res[i][j] = min(res[i][j], res[i][j+1] + 1)
					}
				}
			}
		}

		return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func updateMatrix2(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	queue, visited, res := []int{}, make([][]bool, m),  make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
		res[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				queue = append(queue, i * n + j)
				visited[i][j] = true
			}
		}
	}

	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		x, y := cur / n, cur % n
		for i := 0; i < 4; i++ {
			nx, ny := x + dx[i], y + dy[i]
			if nx < 0 || nx >= m || ny < 0 || ny >=n || visited[nx][ny] {
				continue
			}
			res[nx][ny] = res[x][y] + 1
			queue = append(queue, nx * n + ny)
			visited[nx][ny] = true
		}
	}

	return res
}
// @lc code=end

