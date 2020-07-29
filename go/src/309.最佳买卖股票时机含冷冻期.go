/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	dp := make([][]int, n)
	// dp[i][0]：第i天没有持股且当天没有卖出的最大收益
	// dp[i][1]：第i天持股的最大收益
	// dp[i][2]：第i天没有持股且当天卖出的最大收益
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 3)
	}
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i - 1][0], dp[i - 1][2])
		dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] - prices[i])
		dp[i][2] = dp[i - 1][1] + prices[i]
	}

	return max(dp[n - 1][0], dp[n - 1][2])
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

