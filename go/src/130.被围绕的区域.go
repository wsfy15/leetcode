/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start
var dx []int = []int{1, -1, 0, 0}
var dy []int = []int{0, 0, 1, -1}

func solve(board [][]byte)  {
	// 找到与边界的'0'相连的'0' 这些'O'不会转换为'X'
	if len(board) == 0 {
		return
	}

	m, n := len(board), len(board[0])
	for i := 0; i < n; i++ {
		if board[0][i] == 'O' {
			dfs(board, 0, i)
		}
	}
	for i := 0; i < n; i++ {
		if board[m - 1][i] == 'O' {
			dfs(board, m - 1, i)
		}
	}
	for i := 1; i < m; i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}
	}
	for i := 1; i < m; i++ {
		if board[i][n - 1] == 'O' {
			dfs(board, i, n - 1)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '-' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, row, col int) {
	board[row][col] = '-'
	for i := 0; i < 4; i++ {
		nr, nc := row + dx[i], col + dy[i]
		if nr < 0 || nr >= len(board) || nc < 0 || nc >= len(board[0]) || board[nr][nc] != 'O' {
			continue
		}
		dfs(board, nr, nc)
	}
}
// @lc code=end

