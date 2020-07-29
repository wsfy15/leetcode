/*
 * @lc app=leetcode.cn id=852 lang=golang
 *
 * [852] 山脉数组的峰顶索引
 */

// @lc code=start
func peakIndexInMountainArray(A []int) int {
	n, index := len(A), 1
	for i := 2; i < n - 1; i++ {
		if A[i] > A[index] {
			index = i
		}
	}
	return index
}
// @lc code=end

