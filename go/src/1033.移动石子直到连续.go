/*
 * @lc app=leetcode.cn id=1033 lang=golang
 *
 * [1033] 移动石子直到连续
 */

// @lc code=start
func numMovesStones(a int, b int, c int) []int {
	// 保证 a < b < c
	if b < a {
		a, b = b, a
	}
	if c < b {
		b, c = c, b
	}
	if b < a {
		a, b = b, a
	}

	if b - a == 1 && c - b == 1 {
		return []int{0, 0}
	}
	if b - a <= 2 || c - b <= 2 {
		return []int{1, c - a - 2}
	}
	return []int{2, c - a - 2}
}
// @lc code=end

