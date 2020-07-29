/*
 * @lc app=leetcode.cn id=896 lang=golang
 *
 * [896] 单调数列
 */

// @lc code=start
func isMonotonic(A []int) bool {
	i, n := 1, len(A)
	for i < n && A[i] == A[i - 1] {
		i++
	}
	if i == n {
		return true
	}

	order := -1 // 递减
	if A[i] > A[i - 1] { // 递增
		order = 1
	}

	for i < n {
		if (A[i] - A[i - 1]) * order < 0 {
			return false
		}
		i++
	}

	return true
}
// @lc code=end

