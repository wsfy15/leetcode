/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */

// @lc code=start
// nums[-1] = nums[n] = -∞，
func findPeakElement(nums []int) int {
	n := len(nums)
	low, high := 0, n - 1
	for low < high {
		mid := low + (high - low) >> 1

		// 右侧为升序，那么峰值肯定在右边，而且不可能是mid，所以low=mid+1
		if nums[mid] < nums[mid+1] {
			low = mid + 1
		} else { // 峰值在包含mid的左侧，有可能是mid
			high = mid 
		}
	}

	return low
}
// @lc code=end

