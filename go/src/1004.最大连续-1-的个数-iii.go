/*
 * @lc app=leetcode.cn id=1004 lang=golang
 *
 * [1004] 最大连续1的个数 III
 */

// @lc code=start
func longestOnes(A []int, K int) int {
	n, res, start := len(A), 0, 0
	for i := 0; i < n; i++ {
		if A[i] == 0 {
			if K > 0 {
				K--
			} else {
				res = max(res, i - start)
				for K == 0 && start <= i {
					if A[start] == 0 {
						K++
					}
					start++
				}
				K--
			}
		}
	}

	res = max(res, n - start)
	return res
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}
// @lc code=end

