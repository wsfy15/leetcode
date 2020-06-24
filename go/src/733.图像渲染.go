/*
 * @lc app=leetcode.cn id=733 lang=golang
 *
 * [733] 图像渲染
 */

// @lc code=start
var dx []int = []int{1, -1, 0, 0}
var dy []int = []int{0, 0, 1, -1}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}

	m, n := len(image), len(image[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	dfs(image, visited, sr, sc, newColor)

	return image
}

func dfs(image [][]int, visited [][]bool, row, col, newColor int) {
	visited[row][col] = true
	for i := 0; i < 4; i++ {
		nr, nc := row + dx[i], col + dy[i]
		if nr < 0 || nr >= len(image) || nc < 0 || nc >= len(image[0]) || 
			visited[nr][nc] || image[nr][nc] != image[row][col] {
			continue
		}

		dfs(image, visited, nr, nc, newColor)
	}
	image[row][col] = newColor
}
// @lc code=end

