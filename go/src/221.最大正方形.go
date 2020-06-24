/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m + 1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n + 1)
	}

	maxSide := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i + 1][j + 1] = min(dp[i][j], dp[i][j + 1], dp[i + 1][j]) + 1
				if dp[i + 1][j + 1] > maxSide {
					maxSide = dp[i + 1][j + 1]
				}
			}
		}
	}

	return maxSide * maxSide
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
}
// @lc code=end

