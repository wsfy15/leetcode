/*
 * @lc app=leetcode.cn id=463 lang=golang
 *
 * [463] 岛屿的周长
 */

// @lc code=start
func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res += 4
				// 考虑上和左就行，因为每次遇到会加4条边
				if i > 0 && grid[i - 1][j] == 1 {
					res -= 2 // 将去当前格子的上边和上一个格子的下边
				}
				if j > 0 && grid[i][j - 1] == 1 {
					res -= 2
				}
			}
		}
	}

	return res
}

func islandPerimeter2(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				// 考虑上下左右
				if i == 0 || grid[i - 1][j] == 0 {
					res++
				}
				if i == m - 1 || grid[i + 1][j] == 0 {
					res++
				}
				if j == 0 || grid[i][j - 1] == 0 {
					res++
				}
				if j == n - 1 || grid[i][j + 1] == 0{
					res++
				}
			}
		}
	}

	return res
}
// @lc code=end

