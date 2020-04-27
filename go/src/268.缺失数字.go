/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 缺失数字
 */

// @lc code=start
func missingNumber(nums []int) int {
	size, x, y := len(nums), 0, 0
	for i := 0; i < size; i++ {
		x = x ^ nums[i]
		y = y ^ i
	}
	y = y ^ size

	return x ^ y
}
// @lc code=end

