/*
 * @lc app=leetcode.cn id=1314 lang=golang
 *
 * [1314] 矩阵区域和
 */

// @lc code=start
func matrixBlockSum(mat [][]int, K int) [][]int {
	m, n := len(mat), len(mat[0])
  sum, res := make([][]int, m + 1), make([][]int, m)
	sum[m] = make([]int, n + 1)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		sum[i] = make([]int, n + 1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum[i+1][j+1] = sum[i][j+1] + sum[i+1][j] - sum[i][j] + mat[i][j]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			left, right := max(0, j - K), min(n, j + K + 1)
			up, down := min(m, i + K + 1), max(0, i - K)
			res[i][j] = sum[up][right] + sum[down][left] - sum[down][right] - sum[up][left]
		}
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

