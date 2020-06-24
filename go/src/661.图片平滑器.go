/*
 * @lc app=leetcode.cn id=661 lang=golang
 *
 * [661] 图片平滑器
 */

// @lc code=start
func imageSmoother(M [][]int) [][]int {
	if len(M) == 0 {
		return nil
	}
	
	n, m := len(M), len(M[0])
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			l, r, u, d := max(j - 1, 0), min(j + 1, m - 1), max(i - 1, 0), min(i + 1, n - 1)
			total := 0
			for k := u; k <= d; k++ {
				for h := l; h <= r; h++ {
					total += M[k][h]
				}

			}
			res[i][j] = total / ((r - l + 1) * (d - u + 1))
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

