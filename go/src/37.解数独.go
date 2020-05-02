/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte)  {
	row, col, grid := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
	for i := 0; i < 9; i++ {
		row[i] = make([]bool, 10)
		col[i] = make([]bool, 10)
		grid[i] = make([]bool, 10)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '0'
				row[i][num] = true
				col[j][num] = true
				gridIndex := (i / 3) * 3 + (j / 3)
				grid[gridIndex][num] = true
			}
		}
	}

	solve(board, 0, 0, row, col, grid)
}

func solve(board [][]byte, i, j int, row, col, grid [][]bool) bool {
	if j == 9 {
		i++
		j = 0
	}

	if i == 9 && j == 0 {
		return true
	}

	if board[i][j] == '.' {
		for k := 1; k <= 9; k++ {
			if canPlace(board, i, j, k, row, col, grid) {
				return true
			}
		}
	} else {
		if j == 8 {
			return solve(board, i + 1, 0, row, col, grid)
		} else {
			return solve(board, i, j + 1, row, col, grid)
		}
	}

	return false
}

func canPlace(board [][]byte, i, j, k int, row, col, grid [][]bool) bool {
	gridIndex := (i / 3) * 3 + (j / 3)
	if row[i][k] || col[j][k] || grid[gridIndex][k] {
		return false
	}

	board[i][j] = byte(k + '0')
	row[i][k] = true
	col[j][k] = true
	grid[gridIndex][k] = true

	if solve(board, i, j + 1, row, col, grid) {
		return true
	}

	board[i][j] = '.'
	row[i][k] = false
	col[j][k] = false
	grid[gridIndex][k] = false
	return false
}


// @lc code=end

