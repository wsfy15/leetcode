/*
 * @lc app=leetcode.cn id=377 lang=golang
 *
 * [377] 组合总和 Ⅳ
 */

// @lc code=start
func combinationSum4(nums []int, target int) int {
	sort.Ints(nums)
	if target < nums[0] {
		return 0
	}

	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for j := 0; j < len(nums) && nums[j] <= i; j++ {
			dp[i] += dp[i-nums[j]]
		}
	}

	return dp[target]
}

// @lc code=end

