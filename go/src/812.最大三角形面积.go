/*
 * @lc app=leetcode.cn id=812 lang=golang
 *
 * [812] 最大三角形面积
 */

// @lc code=start
func largestTriangleArea(points [][]int) float64 {
	n, res := len(points), 0.0
	for i := 0; i < n - 2; i++ {
		for j := i + 1; j < n - 1; j++ {
			for k := j + 1; k < n; k++ {
				tmp := area(points, i, j, k)
				res = max(res, tmp)
			}
		}
	}
	return res
}

// http://highscope.ch.ntu.edu.tw/wordpress/?p=66359
func area(points [][]int, i, j, k int) float64 {
	dx1 := points[j][0] - points[i][0]
	dy1 := points[j][1] - points[i][1]
	dx2 := points[k][0] - points[i][0]
	dy2 := points[k][1] - points[i][1]
	tmp := dx1 * dy2 - dx2 * dy1
	if tmp < 0 {
		tmp = -tmp
	}
	return float64(tmp) / 2.0
}

func max(a ...float64) float64 {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}
// @lc code=end

