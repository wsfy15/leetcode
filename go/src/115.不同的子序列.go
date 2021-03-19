/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 */

// @lc code=start
func numDistinct(s string, t string) int {
	m, n := len(t), len(s)
	if m > n {
		return 0
	}

	dp := make([][]int, m + 1) // dp[i][j]: s[:j]中t[:i]的子序列个数
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n + 1)
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i][j - 1]
			if t[i - 1] == s[j - 1] {
				dp[i][j] += dp[i - 1][j - 1]
			}
		}
	}

	return dp[m][n]
}
// @lc code=end

