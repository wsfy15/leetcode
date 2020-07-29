/*
 * @lc app=leetcode.cn id=922 lang=golang
 *
 * [922] 按奇偶排序数组 II
 */

// @lc code=start
func sortArrayByParityII(A []int) []int {
	n := len(A)
	i, j := 1, 0 // i 指向奇数位置， j 指向偶数位置
	for j < n {
		if A[j] & 1 != 0 {
			for i < n && A[i] & 1 == 1 {
				i += 2
			}

			A[i], A[j] = A[j], A[i]
			i += 2
		}

		j += 2
	}

	return A
}
// @lc code=end

