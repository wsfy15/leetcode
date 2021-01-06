/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
func maxProfit(k int, prices []int) int {
	if k == 0 {
		return 0
	}

	n := len(prices)
	if k*2 >= n {
		return greedy(prices)
	}

	dp := make([][][]int, n) // dp[i][j][k]: 第i天，j次买入，k(0:没有持有 1:持有股票) 的利润
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = make([]int, 2)
			dp[i][0][0] = 0
			dp[i][0][1] = math.MinInt32
		}
	}

	for j := 1; j <= k; j++ {
		dp[0][j][0] = 0
		dp[0][j][1] = -prices[0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j-1][0]-prices[i], dp[i-1][j][1])
		}
	}
	return dp[n-1][k][0]
}

func greedy(prices []int) int {
	res := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			res += prices[i+1] - prices[i]
		}
	}
	return res
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

