/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start
func findDuplicate(nums []int) int {
	n := len(nums) - 1
	left, right := 1, n
	for left < right {
		mid := left + (right - left) >> 1 // 中位数
		cnt := 0

		for _, v := range nums {
			if v <= mid {
				cnt++
			}
		}

		if cnt > mid { // [1, mid]范围内的数 的数量大于mid，说明重复在这部分
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
// @lc code=end

