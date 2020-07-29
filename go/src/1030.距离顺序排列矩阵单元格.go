/*
 * @lc app=leetcode.cn id=1030 lang=golang
 *
 * [1030] 距离顺序排列矩阵单元格
 */

// @lc code=start
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	res := [][]int{}
	queue := []int{r0 * 100 + c0}

	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}

	visited := map[int]bool{}
	visited[r0 * 100 + c0] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		x, y := cur / 100, cur % 100
		res = append(res, []int{x, y})

		for i := 0; i < 4; i++ {
			nx, ny := x + dx[i], y + dy[i]
			if !visited[nx*100 + ny] && 0 <= nx && nx < R && 0 <= ny && ny < C {
				queue = append(queue, nx*100 + ny)
				visited[nx*100 + ny] = true
			}
		}
	}

	return res
}
// @lc code=end

