/*
 * @lc app=leetcode.cn id=1006 lang=golang
 *
 * [1006] 笨阶乘
 */

// @lc code=start
func clumsy(N int) int {
	if N == 1 {
		return 1
	} else if N == 2 {
		return 2
	} else if N == 3 {
		return 6
	} else if N == 4 {
		return 7
	}

	if N%4 == 0 {
		return N + 1
	} else if N%4 == 3 {
		return N - 1
	} else {
		return N + 2
	}
}

// @lc code=end

