/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */

// @lc code=start
func sortedSquares(A []int) []int {
	n := len(A) - 1
	res := make([]int, n + 1)
	l, r := 0, n
	for l <= r {
		if abs(A[r]) > abs(A[l]) {
			res[n] = A[r] * A[r]
			r--
		} else {
			res[n] = A[l] * A[l]
			l++
		}
		n--
	}

	return res
}

func abs(x int) int {
	y := x >> 31
	return (x ^ y) - y
}
// @lc code=end

