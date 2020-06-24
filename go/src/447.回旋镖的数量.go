/*
 * @lc app=leetcode.cn id=447 lang=golang
 *
 * [447] 回旋镖的数量
 */

// @lc code=start
func numberOfBoomerangs(points [][]int) int {
	res := 0
	n := len(points)
	for index, point := range points {
		counter := map[int]int{}
		for i := 0; i < n; i++ {
			if i == index {
				continue
			}

			xdiff, ydiff := points[i][0] - point[0], points[i][1] - point[1]
			distance := xdiff * xdiff + ydiff * ydiff
			res += 2 * counter[distance]
			counter[distance]++
		}
	}

	return res
}
// @lc code=end

