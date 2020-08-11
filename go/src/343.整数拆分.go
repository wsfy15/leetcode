/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

// @lc code=start
func integerBreak(n int) int {
	if n < 4 {
		return n - 1
	}

	// 函数y=(n/x)^x (x>0)的最大值，可得x=e时y最大
	// 对于6=2+2+2=3+3来说，3X3 > 2^3
	res := 1
	for n > 4 {
		n -= 3
		res *= 3
	}
	return res * n
}
// @lc code=end

