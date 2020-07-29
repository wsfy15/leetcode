/*
 * @lc app=leetcode.cn id=329 lang=golang
 *
 * [329] 矩阵中的最长递增路径
 */

// @lc code=start
type Point struct {
	x int
	y int
	val int
}

func longestIncreasingPath(matrix [][]int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = 1
		}
	}

	// 先根据每个点的值升序排序，然后依次遍历每一个点
	points := []*Point{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			points = append(points, &Point{
				x: i,
				y: j,
				val: matrix[i][j],
			})
		}
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i].val < points[j].val
	})

	res := 1
	for _, p := range points {
		x, y, val := p.x, p.y, p.val
		if x > 0 && val > matrix[x-1][y] && dp[x-1][y] >= dp[x][y] {
			dp[x][y] = dp[x-1][y] + 1
		}
		if y > 0 && val > matrix[x][y-1] && dp[x][y-1] >= dp[x][y] {
			dp[x][y] = dp[x][y-1] + 1
		}
		if x < n - 1 && val > matrix[x+1][y] && dp[x+1][y] >= dp[x][y] {
			dp[x][y] = dp[x+1][y] + 1
		}
		if y < m - 1 && val > matrix[x][y+1] && dp[x][y+1] >= dp[x][y] {
			dp[x][y] = dp[x][y+1] + 1
		}

		if dp[x][y] > res {
			res = dp[x][y]
		}
	}

	return res
}
// @lc code=end

