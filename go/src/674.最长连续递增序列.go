/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] 最长连续递增序列
 */

// @lc code=start
func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	max, count := 1, 1
	for i := 0; i < n - 1; i++ {
		if nums[i] < nums[i + 1] {
			count++
			if count > max {
				max = count
			}
		} else {
			count = 1
		}
	}
	return max
}
// @lc code=end

