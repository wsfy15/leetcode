/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
var dx []int = []int{1, -1, 0, 0}
var dy []int = []int{0, 0, 1, -1}
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word[0] == board[i][j] && dfs(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, index, i, j int) bool {
	if index == len(word) {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	} 

	if word[index] == board[i][j] {
		board[i][j] = ' '
		for k := 0; k < 4; k++ {
			if dfs(board, word, index + 1, i + dx[k], j + dy[k]) {
				return true
			}
		}
		board[i][j] = word[index]
	}
	 
	return false
}

// @lc code=end

