/*
 * @lc app=leetcode.cn id=201 lang=golang
 *
 * [201] 数字范围按位与
 */

// @lc code=start
func rangeBitwiseAnd(m int, n int) int {
	// 最多只有高几位能保持为1
	offset := 0
	for m != n {
		m >>= 1
		n >>= 1
		offset++
	}
	return m << offset
}
// @lc code=end

