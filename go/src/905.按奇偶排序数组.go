/*
 * @lc app=leetcode.cn id=905 lang=golang
 *
 * [905] 按奇偶排序数组
 */

// @lc code=start
func sortArrayByParity2(A []int) []int {
	sort.Slice(A, func(i, j int) bool {
		if A[i] % 2 == 0 {
			return true
		} else if A[j] % 2 == 0 {
			return false
		}
		return true
	})
	return A
}

func sortArrayByParity(A []int) []int {
	n := len(A)
	if n < 2 {
		return A
	}

	i, j := 0, 0
	for j < n {
		if A[j] % 2 == 0 {
			A[i], A[j] = A[j], A[i]
			i++
		}
		j++
	}

	return A
}
// @lc code=end

