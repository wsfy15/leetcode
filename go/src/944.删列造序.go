/*
 * @lc app=leetcode.cn id=944 lang=golang
 *
 * [944] 删列造序
 */

// @lc code=start
func minDeletionSize(A []string) int {
	n, m := len(A), len(A[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			if A[j-1][i] > A[j][i] {
				res++
				break
			}
		}
	}
	return res
}
// @lc code=end

