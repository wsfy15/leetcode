/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
// [1] 1
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	low, high := 0, len(nums) - 1
	res := []int{-1, -1}

	// 找左边界
	for low < high {
		mid := low + (high - low) >> 1

		if nums[mid] >= target {
			high = mid
		} else {
			low = mid + 1
		}
	}

	if nums[low] == target {
		res[0] = low

		// 找右边界 找到比target大的下标，然后-1
		high = len(nums)
		low++
		for low < high {
			mid := low + (high - low) >> 1

			if nums[mid] > target {
				high = mid
			} else {
				low = mid + 1
			}
		}
		res[1] = low - 1
	}

	return res
}
// @lc code=end

