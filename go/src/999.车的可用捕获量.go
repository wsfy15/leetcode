/*
 * @lc app=leetcode.cn id=999 lang=golang
 *
 * [999] 车的可用捕获量
 */

// @lc code=start
func numRookCaptures(board [][]byte) int {
	res := 0
	x, y := -1, -1 // 车所在位置
	for i := 0;  i < 8; i++ {
		flag := false
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				x, y = i, j
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}

	for i := x + 1; i < 8; i++ {
		if board[i][y] == 'p' {
			res++
			break
		} else if board[i][y] == 'B' {
			break
		}
	}

	for i := x - 1; i >= 0; i-- {
		if board[i][y] == 'p' {
			res++
			break
		} else if board[i][y] == 'B' {
			break
		}
	}

	for j := y + 1; j < 8; j++ {
		if board[x][j] == 'p' {
			res++
			break
		} else if board[x][j] == 'B' {
			break
		}
	}

	for j := y - 1; j >= 0; j-- {
		if board[x][j] == 'p' {
			res++
			break
		} else if board[x][j] == 'B' {
			break
		}
	}

	return res
}
// @lc code=end

