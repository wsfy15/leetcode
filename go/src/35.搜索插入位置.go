/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
// 找到大于等于target的第一个位置
func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)
	for low < high {
		mid := low + (high - low) >> 1

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}
// @lc code=end

