/*
 * @lc app=leetcode.cn id=375 lang=golang
 *
 * [375] 猜数字大小 II
 */

// @lc code=start
func getMoneyAmount(n int) int {
	dp := make([][]int, n + 1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n + 1)
	}

	// for i := 1; i <= n; i++ { // 只有1个数字，开销为0
	// 	dp[i][i] = 0
	// }

	for k := 2; k <= n; k++ { // k个数
		for i := 1; i <= n - k + 1; i++ {
			// 依次选择[i, i + k - 1]内的每个数，考察左右两部分的值
			dp[i][i + k - 1] = min(i + dp[i+1][i+k-1], dp[i][i+k-2] + i + k - 1)
			for j := i + 1; j < i + k - 1; j++ {
				dp[i][i + k - 1] = min(dp[i][i + k - 1], max(dp[i][j-1], dp[j+1][i+k-1]) + j)
			} 
		}
	}

	return dp[1][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

