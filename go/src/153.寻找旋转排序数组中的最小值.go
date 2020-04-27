/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	low, high := 0, len(nums) - 1
	for low < high {
		mid := low + (high - low) >> 1

		if nums[mid] > nums[high] { // 最小值位于右边
			low = mid + 1
		} else {
			high = mid
		}
	}

	return nums[low]
}
// @lc code=end

