/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	// 将n放到nums[n - 1]的位置  n < length
	var length = len(nums)
	for i := 0; i < length; i++ {
		for nums[i] > 0 && nums[i] <= length && nums[nums[i]-1] != nums[i] {
			var tmp = nums[i]
			nums[i] = nums[nums[i]-1]
			nums[tmp-1] = tmp
		}
	}

	for i := 0; i < length; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return length + 1
}

// @lc code=end
