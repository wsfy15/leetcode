/*
 * @lc app=leetcode.cn id=840 lang=golang
 *
 * [840] 矩阵中的幻方
 */

// @lc code=start
func numMagicSquaresInside(grid [][]int) int {
	// 1 到 9 不同 数字
	row, col := len(grid), len(grid[0])
	if row < 3 || col < 3 {
		return 0
	}

	res := 0
	for i := 0; i < row - 2; i++ {
		for j := 0; j < col - 2; j++ {
			if check(grid, i, j) {
				res++
			}
		}
	}

	return res
}

func check(grid [][]int, row, col int) bool {
	same := true
	for i := row; i < row + 3; i++ {
		for j := col; j < col + 3; j++ {
			if grid[i][j] < 1 || grid[i][j] > 9 {
				return false
			}
			if grid[i][j] != grid[row][col] {
				same = false
			}
		}
	}

	if same {
		return false
	}

	num := grid[row][col] + grid[row][col + 1] + grid[row][col + 2]
	if grid[row][col] + grid[row+1][col+1] + grid[row+2][col+2] != num {
		return false
	}
	if grid[row][col+2] + grid[row+1][col+1] + grid[row+2][col] != num {
		return false
	}

	for i := row + 1; i < row + 3; i++ {
		if grid[i][col] + grid[i][col+1] + grid[i][col+2] != num {
			return false
		}
	}

	for j := col; j < col + 3; j++ {
		if grid[row][j] + grid[row+1][j] + grid[row+2][j] != num {
			return false
		}
	}

	return true
}
// @lc code=end

