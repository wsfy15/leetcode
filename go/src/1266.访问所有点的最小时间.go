/*
 * @lc app=leetcode.cn id=1266 lang=golang
 *
 * [1266] 访问所有点的最小时间
 */

// @lc code=start
func minTimeToVisitAllPoints(points [][]int) int {
	n, res := len(points), 0
	for i := 1; i < n; i++ {
		res += max(abs(points[i][1]-points[i-1][1]), abs(points[i][0]-points[i-1][0]))
	}

	return res
}

func abs(x int) int {
	y := x >> 31
	return (x ^ y) - y
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

