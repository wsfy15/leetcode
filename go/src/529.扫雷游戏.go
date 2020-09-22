/*
 * @lc app=leetcode.cn id=529 lang=golang
 *
 * [529] 扫雷游戏
 */

// @lc code=start
var dx []int = []int{0, 0, 1, -1, 1, 1, -1, -1}
var dy []int = []int{1, -1, 0, 0, 1, -1, 1, -1}

func updateBoard(board [][]byte, click []int) [][]byte {
	dfs(board, click[0], click[1])
	return board
}

func dfs(board [][]byte, row, col int) {
	m, n := len(board), len(board[0])
	if row < 0 || row >= m || col < 0 || col >= n {
		return
	}
	
	if board[row][col] == 'M' {
		board[row][col] = 'X'
	} else if board[row][col] == 'E' {
		count := getCount(board, row, col)
		if count == 0 {
			board[row][col] = 'B'
			for i := 0; i < 8; i++ {
				dfs(board, row + dx[i], col + dy[i])
			}
		} else {
			board[row][col] = '0' + byte(count)
		}
	}
}

func getCount(board [][]byte, row, col int) int {
	m, n := len(board), len(board[0])
	count := 0
	for i := 0; i < 8; i++ {
		x, y := row + dx[i], col + dy[i]
		if x < 0 || x >= m || y < 0 || y >= n {
			continue
		}

		if board[x][y] == 'M' {
			count++
		}
	}

	return count
}
// @lc code=end

