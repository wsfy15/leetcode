/*
 * @lc app=leetcode.cn id=363 lang=golang
 *
 * [363] 矩形区域不超过 K 的最大数值和
 */

// @lc code=start
func maxSumSubmatrix(matrix [][]int, k int) int {
	n, m, res := len(matrix), len(matrix[0]), math.MinInt32

	// 固定左右边界，求从上往下行的前缀和
	for left := 0; left < m; left++ {
		rowSum := make([]int, n+1)
		for right := left; right < m; right++ {
			for i := 1; i <= n; i++ {
				rowSum[i] += matrix[i-1][right]
			}

			res = max(res, getMax(rowSum, k))
		}
	}

	return res
}

func getMax(sum []int, k int) int {
	res, n := math.MinInt32, len(sum)
	for i := 1; i < n; i++ {
		cur := 0
		for j := i; j < n; j++ {
			cur += sum[j]
			if cur > res && cur <= k {
				res = cur
			}
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

