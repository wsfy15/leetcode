/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 */

// @lc code=start
func trailingZeroes(n int) int {
	res, divisor := 0, 5

	for n >= divisor {
		res += n / divisor
		divisor *= 5
	}

	return res
}
// @lc code=end

