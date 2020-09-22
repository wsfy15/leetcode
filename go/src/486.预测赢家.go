/*
 * @lc app=leetcode.cn id=486 lang=golang
 *
 * [486] 预测赢家
 */

// @lc code=start
func PredictTheWinner(nums []int) bool {
	n := len(nums)
	if n%2 == 0 { // 偶数个玩家1必胜，因为若玩家1的玩法输了，由于它是先手，它可以换成玩家2的玩法
		return true
	}

	// dp[i][j] 表示 从nums[i]到nums[j]先手比另一位玩家多的最大分数
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}

	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			dp[j][i] = max(nums[i]-dp[j][i-1], nums[j]-dp[j+1][i])
		}
	}

	return dp[0][n-1] >= 0
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

