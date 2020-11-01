/*
 * @lc app=leetcode.cn id=845 lang=golang
 *
 * [845] 数组中的最长山脉
 */

// @lc code=start
func longestMountain(A []int) int {
	n := len(A)
	if n < 3 {
		return 0
	}

	res := 0
	for i := 1; i < n-1; i++ {
		if A[i-1] < A[i] && A[i] > A[i+1] {
			l, r := i-1, i+1
			for l > 0 && A[l] > A[l-1] {
				l--
			}
			for r < n-1 && A[r] > A[r+1] {
				r++
			}
			res = max(res, r-l+1)
		}
	}

	return res
}

func max(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v > ans {
			ans = v
		}
	}
	return ans
}

// @lc code=end

