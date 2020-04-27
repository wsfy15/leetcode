/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	
	x, y := 0, nums[0]
	for i := 1; i < size; i++ {
		x, y = max(x, y), x + nums[i]
	}
	return max(x, y)

	// dp := make([][]int, size)
	// for i := 0; i < size; i++ {
	// 	dp[i] = make([]int, 2) // 0表示不偷 1表示偷
	// }

	// dp[0][0], dp[0][1] = 0, nums[0]
	// for i := 1; i < size; i++ {
	// 	dp[i][0] = max(dp[i-1][0], dp[i-1][1])
	// 	dp[i][1] = dp[i-1][0] + nums[i]
	// }

	// return max(dp[size-1][0], dp[size-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

