/*
 * @lc app=leetcode.cn id=1014 lang=golang
 *
 * [1014] 最佳观光组合
 */

// @lc code=start
func maxScoreSightseeingPair(A []int) int {
	// preMax记录最大的A[i] + i
	preMax, res := A[0] + 0, 0
	n := len(A)
	for i := 1; i < n; i++ {
		res = max(res, preMax + A[i] - i)
		preMax = max(preMax, i + A[i])
	}
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

