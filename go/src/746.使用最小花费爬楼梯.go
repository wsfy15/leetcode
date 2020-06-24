/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n + 1)
	dfs(cost, dp, n)
	return dp[n]
}

func dfs(cost, dp []int, index int) int {
	if index < 2 {
		return 0
	}
	if dp[index] > 0 {
		return dp[index]
	}

	dp[index] = min(cost[index - 1] + dfs(cost, dp, index - 1), cost[index - 2] + dfs(cost, dp, index - 2))
	return dp[index]
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res > v {
			res = v
		}
	}
	return res
}
// @lc code=end

