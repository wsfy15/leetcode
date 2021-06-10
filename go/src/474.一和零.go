/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
	sLen := len(strs)
	// dp[i][j][k]: 考虑strs[:i]，最多j个0和k个1，字符串的最大数量
	dp := make([][][]int, sLen+1)
	for i := 0; i <= sLen; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}

	for i, str := range strs {
		zero, one := getZeroAndOne(str)
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[i+1][j][k] = dp[i][j][k]
				if j >= zero && k >= one {
					dp[i+1][j][k] = max(dp[i+1][j][k], dp[i][j-zero][k-one]+1)
				}
			}
		}
	}

	return dp[sLen][m][n]
}

func getZeroAndOne(s string) (zero int, one int) {
	for _, ch := range s {
		if ch == '1' {
			one++
		} else {
			zero++
		}
	}
	return
}

func max(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v > ans {
			ans = v
		}
	}
	return ans
}

// @lc code=end

