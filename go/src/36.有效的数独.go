/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	horizon := [9]map[int]bool{}
	vertical := [9]map[int]bool{}
	square := [9]map[int]bool{}
	for i := 0; i < 9; i++ {
		horizon[i] = make(map[int]bool)
		vertical[i] = make(map[int]bool)
		square[i] = make(map[int]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			num := int(board[i][j] - '0')
			if val, ok := horizon[i][num]; ok && val {
				return false
			}
			if val, ok := vertical[j][num]; ok && val {
				return false
			}
			if val, ok := square[(i/3)*3 + j / 3][num]; ok && val {
				return false
			}
			horizon[i][num] = true
			vertical[j][num] = true
			square[(i/3)*3 + j / 3][num] = true
		}
	}
	return true
}
// @lc code=end

