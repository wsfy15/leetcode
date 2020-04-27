/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] 交错字符串
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) + len(s2) != len(s3) {
		return false
	}
	m, n := len(s1), len(s2)
	dp := make([][]bool, m + 1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n + 1)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && s2[j - 1] == s3[j - 1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s1[i - 1] == s3[i - 1]
			} else {
				dp[i][j] = (dp[i][j-1] && s2[j - 1] == s3[i + j - 1]) || (dp[i-1][j] && s1[i - 1] == s3[i + j - 1])
			}
		}
	}

	return dp[m][n]
}

func isInterleave2(s1 string, s2 string, s3 string) bool {
	size := len(s3)
	if len(s1) + len(s2) != size {
		return false
	}

	return dfs(s1, 0, s2, 0, s3, 0)
}

func dfs(s1 string, i int, s2 string, j int, s3 string, k int) bool {
	if k == len(s3) {
		return true
	}

	flag := false
	if i < len(s1) && s1[i] == s3[k] {
		flag = dfs(s1, i + 1, s2, j, s3, k + 1)
	}

	if flag {
		return true
	}

	if j < len(s2) && s2[j] == s3[k] {
		flag = dfs(s1, i, s2, j + 1, s3, k + 1)
	}

	return flag
}
// @lc code=end

