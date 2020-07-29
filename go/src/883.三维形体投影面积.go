/*
 * @lc app=leetcode.cn id=883 lang=golang
 *
 * [883] 三维形体投影面积
 */

// @lc code=start
func projectionArea(grid [][]int) int {
	n := len(grid)
	res := 0

	// 记录从前面（每列）和侧面（每行）看到的面积
	front, side := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				res++
				front[j] = max(front[j], grid[i][j])
				side[i] = max(side[i], grid[i][j])
			}
		}
	}

	for i := 0; i < n; i++ {
		res += front[i] + side[i]
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

