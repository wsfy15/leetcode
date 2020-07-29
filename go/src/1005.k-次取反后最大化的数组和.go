/*
 * @lc app=leetcode.cn id=1005 lang=golang
 *
 * [1005] K 次取反后最大化的数组和
 */

// @lc code=start
func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)
	n := len(A)
	for i := 0; i < n; i++ {
		if A[i] < 0 && K > 0 {
			A[i] = -A[i]
			K--
		} else {
			break
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		res += A[i]
	}

	if K > 0 && K & 1 == 1 {
		res -= min(A...) * 2
	}

	return res
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res > v {
			res = v
		}
	}
	return res
}
// @lc code=end

