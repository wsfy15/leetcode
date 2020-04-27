/*
 * @lc app=leetcode.cn id=289 lang=golang
 *
 * [289] 生命游戏
 */

// @lc code=start
func gameOfLife(board [][]int)  {
	// 用二进制的第二位表示下一个状态
	// 10：当前为dead，下一时刻为live
	// 11：当前为live，下一时刻为live

	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])

	dx := []int{0, 0, 1, 1, 1, -1, -1, -1}
	dy := []int{-1, 1, 1, 0, -1, 1, 0, -1}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count := 0
			for k := 0; k < 8; k++ {
				x, y := i + dx[k], j + dy[k]
				if x < 0 || x == m || y < 0 || y == n {
					continue
				}
				if board[x][y] & 1 == 1 {
					count++
				}
			}

			if count == 3 {
				board[i][j] += 2
			} else if count == 2 && board[i][j] == 1 {
				board[i][j] += 2
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] >>= 1
		}
	}
}
// @lc code=end

