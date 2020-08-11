/*
 * @lc app=leetcode.cn id=1137 lang=golang
 *
 * [1137] 第 N 个泰波那契数
 */

// @lc code=start
func tribonacci(n int) int {
	if n < 2 {
		return n
	}

	a, b, c := 0, 1, 1
	for i := 2; i < n; i++ {
		a, b, c = b, c, a+b+c
	}
	return c
}

// @lc code=end

