/*
 * @lc app=leetcode.cn id=961 lang=golang
 *
 * [961] 重复 N 次的元素
 */

// @lc code=start
func repeatedNTimes(A []int) int {
	n := len(A)
	for i := 0; i < n - 2; i++ {
		if A[i] == A[i+1] || A[i] == A[i+2] {
			return A[i]
		}
	}

	return A[n-1]
}
// @lc code=end

