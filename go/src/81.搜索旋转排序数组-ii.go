/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
func search(nums []int, target int) bool {
	low, high := 0, len(nums) - 1
	if high == - 1 {
		return false
	}

	for low < high {
		for low + 1 < high && nums[low] == nums[low + 1] {
			low++
		}
		for low + 1 < high && nums[high] == nums[high - 1] {
			high--
		}

		mid := low + (high - low) >> 1

		if nums[mid] == target {
			return true
		}

		if nums[low] <= nums[mid] {
			if nums[low] <= target && target <= nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return nums[low] == target
}
// @lc code=end

