/*
 * @lc app=leetcode.cn id=976 lang=golang
 *
 * [976] 三角形的最大周长
 */

// @lc code=start
func largestPerimeter(A []int) int {
	sort.Ints(A)
	n := len(A)
	for i := n - 1; i >= 2; i-- {
		if A[i-1] + A[i-2] > A[i] {
			return A[i] + A[i-1] + A[i-2]
		}
	}
	return 0
}
// @lc code=end

