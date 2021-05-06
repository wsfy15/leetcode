/*
 * @lc app=leetcode.cn id=740 lang=golang
 *
 * [740] 删除与获得点数
 */

// @lc code=start
func deleteAndEarn(nums []int) int {
	vals, dp := make([]int, 1e4+1), make([]int, 1e4+1)
	for _, num := range nums {
		vals[num] += num
	}
	dp[0], dp[1] = 0, vals[1]

	for i := 2; i <= 1e4; i++ {
		dp[i] = max(dp[i-1], vals[i]+dp[i-2])
	}

	return dp[1e4]
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

