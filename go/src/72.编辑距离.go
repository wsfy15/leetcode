/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	states := make([][]int, m) // states[i][j]表示word1[0...i] 与 word2[0...j]的编辑距离
	for i := range states {
		states[i] = make([]int ,n)
	}
	
	for i := 0; i < n; i++ { // 比较word1[0] 与 word2的距离
		if word1[0] == word2[i] {
			states[0][i] = i
		} else if i > 0 {
			states[0][i] = states[0][i - 1] + 1
		} else {
			states[0][0] = 1
		}
	}

	for i := 1; i < m; i++ { // 比较word2[0] 与 word1的距离
		if word1[i] == word2[0] {
			states[i][0] = i
		} else {
			states[i][0] = states[i - 1][0] + 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i] == word2[j] {
				states[i][j] = min(states[i - 1][j - 1], states[i - 1][j] + 1, states[i][j - 1] + 1)
			} else {
				states[i][j] = 1 + min(states[i - 1][j - 1], states[i - 1][j], states[i][j - 1])
			}
		}
	}

	return states[m - 1][n - 1]
}

func min(a, b, c int) int {
	res := a
	if res > b {
		res = b
	}
	if res > c {
		res = c
	}
	return res
}
// @lc code=end

