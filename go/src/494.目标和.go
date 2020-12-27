/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
var res int

func findTargetSumWays(nums []int, S int) int {
	res = 0
	dfs(nums, S, 0)

	return res
}

func dfs(nums []int, target, index int) {
	if index == len(nums) {
		if target == 0 {
			res++
		}
		return
	}

	dfs(nums, target-nums[index], index+1)
	dfs(nums, target+nums[index], index+1)
}

// @lc code=end

