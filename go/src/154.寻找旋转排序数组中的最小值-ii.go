/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

// @lc code=start
func findMin(nums []int) int {
	left, right := 0, len(nums) - 1
	for left < right {
		mid := left + (right - left) >> 1

		if nums[mid] > nums[right] { // 肯定在右边
			left = mid + 1
		} else if nums[mid] < nums[right] { // 肯定在左边
			right = mid
		} else { // [1,1,0,1] [1,0,1,1,1] 既有可能在左边，也有可能在右边，直接right--可以保证不会越界，也不会丢失最小值
			right--
		}
	}

	return nums[left]
}
// @lc code=end

