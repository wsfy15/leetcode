/*
 * @lc app=leetcode.cn id=879 lang=golang
 *
 * [879] 盈利计划
 */

// @lc code=start
func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	var mod int = 1e9 + 7
	sum := 0
	for _, p := range profit {
		sum += p
	}

	dp := make([][]int, n+1) // dp[i][j]: i个人产生j收益的方法数
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, sum+1)
	}

	dp[0][0] = 1
	for i := 0; i < len(group); i++ {
		g, p := group[i], profit[i]
		for j := n; j >= g; j-- {
			for k := sum; k >= p; k-- {
				dp[j][k] = (dp[j][k] + dp[j-g][k-p]) % mod
			}
		}
	}

	res := 0
	for i := 0; i <= n; i++ {
		for j := minProfit; j <= sum; j++ {
			res = (res + dp[i][j]) % mod
		}
	}
	return res
}

// @lc code=end

