/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	end := n * n
	num := 1
	i, j, direction := 0, 0, 0
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	for num <= end {
		res[i][j] = num
		num++

		ni, nj := i + dx[direction], j + dy[direction]
		if ni >= n || ni < 0 || nj >= n || nj < 0 || res[ni][nj] > 0 {
			direction = (direction + 1) % 4
			ni, nj = i + dx[direction], j + dy[direction]
		}

		i, j = ni, nj
	}
	return res
}
// @lc code=end

