/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] 用最少数量的箭引爆气球
 */

// @lc code=start
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i][1] == points[j][1] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})

	res, right := 1, points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > right {
			res++
			right = points[i][1]
		}
	}

	return res
}

// @lc code=end

