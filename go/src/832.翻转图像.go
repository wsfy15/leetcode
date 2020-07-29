/*
 * @lc app=leetcode.cn id=832 lang=golang
 *
 * [832] 翻转图像
 */

// @lc code=start
func flipAndInvertImage(A [][]int) [][]int {
	n, m := len(A), len(A[0])
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
		for j := 0; j < m; j++ {
			res[i][j] = (A[i][m - j - 1] + 1) % 2
		}
	}

	return res
}
// @lc code=end

