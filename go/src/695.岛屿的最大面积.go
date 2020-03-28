/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
var max int
var cur int
var dx []int = []int{1, -1, 0, 0}
var dy []int = []int{0, 0, 1, -1}
func maxAreaOfIsland(grid [][]int) int {
	// 并查集，最大的集合
	// 染色法
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	max = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				cur = 0 
				util(grid, i, j)
			}
		}
	}

	return max
}

func util(grid [][]int, i, j int) {
	cur++
	if cur > max {
		max = cur
	}
	grid[i][j] = 0
	for k := 0; k < 4; k++ {
		x := i + dx[k]
		y := j + dy[k]
		if 0 <= x && x < len(grid) && 0 <= y && y < len(grid[0]) && grid[x][y] == 1 {
			util(grid, x, y)
		}
	}
}
// @lc code=end

