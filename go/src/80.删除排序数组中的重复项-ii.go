/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除排序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	left, right := 1, 1
	for right < n {
		if nums[right] == nums[right - 1] {
			for right + 1 < n && nums[right] == nums[right + 1] {
				right++
			}
		} 
		nums[left] = nums[right]
		left++
		right++
	}

	return left
}
// @lc code=end

