/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	low, high := 0, len(nums) - 1
	for low < high { // 因为只有low=mid+1，而high=mid，high肯定不满足条件，所以可以用low < high，而不是<=
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

	if nums[low] != target {
		return -1
	}
	return low
}
// @lc code=end

