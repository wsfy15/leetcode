/*
 * @lc app=leetcode.cn id=476 lang=golang
 *
 * [476] 数字的补数
 */

// @lc code=start
func findComplement(num int) int {
	res, factor := 0, 1
	for num > 0 {
		if num & 1 == 0 {
			res |= factor
		}
		factor <<= 1
		num >>= 1
	}
	return res
}
// @lc code=end

