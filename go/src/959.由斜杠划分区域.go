/*
 * @lc app=leetcode.cn id=959 lang=golang
 *
 * [959] 由斜杠划分区域
 */

// @lc code=start
var dx []int = []int{0, 0, 1, -1}
var dy []int = []int{1, -1, 0, 0}
func regionsBySlashes(grid []string) int {
	n := len(grid)
	// 每个 1X1 方格转化为3X3方格，用岛屿个数思路求解
	g := make([][]int, n * 3)
	for i := 0; i < n * 3; i++ {
		g[i] = make([]int, n * 3)
	}

	for i, str := range grid {
		for j, ch := range []byte(str) {
			switch ch {
			case '/':
				g[i * 3][j * 3 + 2] = 1
				g[i * 3 + 1][j * 3 + 1] = 1
				g[i * 3 + 2][j * 3] = 1
			case '\\':
				g[i * 3][j * 3] =	1
				g[i * 3 + 1][j * 3 + 1] = 1
				g[i * 3 + 2][j * 3 + 2] = 1
			}
		}
	}

	res := 0
	for i := 0; i < n * 3; i++ {
		for j := 0; j < n * 3; j++ {
			if g[i][j] == 0 {
				res++
				dfs(i, j, g)
			}
		}
	}

	return res
}

func dfs(x, y int, g [][]int) {
	g[x][y] = 1
	for i := 0; i < 4; i++ {
		nx, ny := x + dx[i], y + dy[i]
		if 0 <= nx && nx < len(g) && 0 <= ny && ny < len(g) && g[nx][ny] == 0 {
			dfs(nx, ny, g)
		}
	}
}
// @lc code=end

