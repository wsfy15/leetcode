/*
 * @lc app=leetcode.cn id=581 lang=golang
 *
 * [581] 最短无序连续子数组
 */

// @lc code=start
func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	// 从左往右找到最后一个破坏升序的元素
	max, high := nums[0], 0
	for i := 1; i < n; i++ {
		if nums[i] >= max {
			max = nums[i]
		} else {
			high = i
		}
	}

	// 从右往左找到最后一个破坏降序的元素
	min, low := nums[n - 1], n - 1
	for i := n - 2; i >= 0; i-- {
		if nums[i] <= min {
			min = nums[i]
		} else {
			low = i
		}
	}

	if high > low {
		return high - low + 1
	}
	return 0
}
// @lc code=end

