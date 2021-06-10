/*
 * @lc app=leetcode.cn id=1049 lang=golang
 *
 * [1049] 最后一块石头的重量 II
 */

// @lc code=start
func lastStoneWeightII(stones []int) int {
	// 相当于将一堆石头分成两堆，使两堆的重量接近，重量差最小值就是答案
	sum := 0
	for _, stone := range stones {
		sum += stone
	}

	// 求小于sum/2 但最大重量的石堆的重量
	target := sum / 2
	dp := make([]int, target+1)
	for _, stone := range stones {
		for i := target; i >= stone; i-- {
			dp[i] = max(dp[i], dp[i-stone]+stone)
		}
	}

	return sum - 2*dp[target]
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

