/*
 * @lc app=leetcode.cn id=892 lang=golang
 *
 * [892] 三维形体的表面积
 */

// @lc code=start
func surfaceArea(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	res := 0
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, -1, 1}

	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if grid[i][j] > 0 {
				res += 2
				for k := 0; k < 4; k++ {
					x, y := i + dx[k], j + dy[k]
					if x < 0 || x == rows || y < 0 || y == rows {
						res += grid[i][j]
					} else if grid[i][j] > grid[x][y] {
						res += grid[i][j] - grid[x][y]
					}
				}
			}
		}
	}

	return res
}
// @lc code=end

