/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] 通配符匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	if n == 0 {
		return m == 0
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			if p[i] != '*' {
				return false
			}
		}
		return true
	}

	dp := make([][]bool, m + 1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n + 1)
	}

	dp[0][0] = true
	j := 0
	for j < n && p[j] == '*' { // 处理pattern开头的连续*号
		j++
		for i := 0; i <= m; i++ {
			dp[i][j] = true
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if ((dp[j - 1][i] || dp[j][i - 1]) && p[i - 1] == '*') ||
				(dp[j - 1][i - 1] && (p[i - 1] == s[j - 1] || p[i - 1] == '?' )) {
				dp[j][i] = true
			}
		}
	}

	return dp[m][n]
}
// @lc code=end

