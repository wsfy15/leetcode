/*
 * @lc app=leetcode.cn id=941 lang=golang
 *
 * [941] 有效的山脉数组
 */

// @lc code=start
func validMountainArray(A []int) bool {
	n := len(A)
	if n < 3 {
		return false
	}

	i := 0
	for i < n - 1 {
		if A[i] < A[i + 1] {
			i++
		}  else if A[i] == A[i + 1] {
			return false
		} else {
			break
		}
	}

	if i == n - 1 || i == 0 {
		return false
	}

	for i < n - 1 {
		if A[i] > A[i + 1] {
			i++
		} else {
			return false
		}
	}

	return true
}
// @lc code=end

