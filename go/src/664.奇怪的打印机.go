/*
 * @lc app=leetcode.cn id=664 lang=golang
 *
 * [664] 奇怪的打印机
 */

// @lc code=start
func strangePrinter(s string) int {
	n := len(s)
	// dp[i][j]: s[i:j+1] 的最少打印次数
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		dp[i][i] = 1
	}
	dp[n][n] = 0

	for l := 2; l <= n; l++ { // 长度
		for i := 0; i <= n-l; i++ { // 起点
			j := i + l - 1
			dp[i][j] = dp[i+1][j] + 1
			for k := i + 1; k <= j; k++ { // 分割点
				if s[i] == s[k] {
					dp[i][j] = min(dp[i][j], dp[i][k-1]+dp[k+1][j])
				}
			}
		}
	}

	return dp[0][n-1]
}

func min(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v < ans {
			ans = v
		}
	}
	return ans
}

// @lc code=end

