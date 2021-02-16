/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	if m * n == 0 {
		return 0
	}

	dp := make([][]int, n + 1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m + 1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			} else {
				dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])
			}
		}
	}

	return dp[n][m]
}

func longestCommonSubsequence2(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	if m * n == 0 {
		return 0
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		if text1[i] == text2[0] || (i > 0 && dp[i - 1][0] == 1) {
			dp[i][0] = 1
		}
	}
	for i := 0; i < m; i++ {
		if text1[0] == text2[i] || (i > 0 && dp[0][i - 1] == 1) {
			dp[0][i] = 1
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			} else {
				dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])
			}
		}
	}

	return dp[n-1][m-1]
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}
// @lc code=end

