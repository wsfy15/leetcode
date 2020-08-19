/*
 * @lc app=leetcode.cn id=1009 lang=golang
 *
 * [1009] 十进制整数的反码
 */

// @lc code=start
func bitwiseComplement(N int) int {
	num := 1
	for num < N {
		num = (num << 1) + 1
	}
	return num ^ N
}

// @lc code=end

