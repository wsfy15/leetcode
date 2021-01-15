/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	dp := make([][][]int, n) // dp[i][j][k]:第i天，j(0:未持股 1:持股) k次交易 的利润
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, 2)
		dp[i][0], dp[i][1] = make([]int, 3), make([]int, 3)
	}
	dp[0][1][0] = -prices[0]
	dp[0][0][1], dp[0][0][2] = math.MinInt32, math.MinInt32
	dp[0][1][1], dp[0][1][2] = math.MinInt32, math.MinInt32

	for i := 1; i < n; i++ {
		dp[i][0][0] = 0
		dp[i][0][1] = max(dp[i-1][0][1], dp[i-1][1][0]+prices[i])
		dp[i][0][2] = max(dp[i-1][0][2], dp[i-1][1][1]+prices[i])
		dp[i][1][0] = max(dp[i-1][0][0]-prices[i], dp[i-1][1][0])
		dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][1]-prices[i])
		dp[i][1][2] = math.MinInt32
	}

	return max(dp[n-1][0][2], dp[n-1][0][1], 0)
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

