/*
 * @lc app=leetcode.cn id=973 lang=golang
 *
 * [973] 最接近原点的 K 个点
 */

// @lc code=start
type point struct {
	x, y     int
	distance int
}

func kClosest(points [][]int, K int) [][]int {
	n := len(points)
	p := make([]point, n)
	for i := 0; i < n; i++ {
		x, y := points[i][0], points[i][1]
		p[i] = point{
			x:        x,
			y:        y,
			distance: x*x + y*y,
		}
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i].distance < p[j].distance
	})

	res := make([][]int, K)
	for i := 0; i < K; i++ {
		res[i] = []int{p[i].x, p[i].y}
	}
	return res
}

// @lc code=end

