/*
 * @lc app=leetcode.cn id=1009 lang=golang
 *
 * [1009] 十进制整数的反码
 */

// @lc code=start
func bitwiseComplement(N int) int {
	tmp := N
	num := 1
	for tmp > 0 {
		tmp >>= 1
		num <<= 1
	}
	return num - N - 1
}
// @lc code=end

