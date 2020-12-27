/*
 * @lc app=leetcode.cn id=417 lang=golang
 *
 * [417] 太平洋大西洋水流问题
 */

// @lc code=start
var dx []int = []int{1, -1, 0, 0}
var dy []int = []int{0, 0, 1, -1}

func pacificAtlantic(matrix [][]int) [][]int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])

	canReachP, canReachA := make([][]bool, m), make([][]bool, m)
	for i := 0; i < m; i++ {
		canReachP[i] = make([]bool, n)
		canReachA[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		dfs(canReachP, matrix, i, 0)
		dfs(canReachA, matrix, i, n-1)
	}

	for i := 0; i < n; i++ {
		dfs(canReachP, matrix, 0, i)
		dfs(canReachA, matrix, m-1, i)
	}

	res := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if canReachP[i][j] && canReachA[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}

func dfs(reach [][]bool, matrix [][]int, x, y int) {
	reach[x][y] = true
	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if nx < 0 || nx >= len(reach) || ny < 0 || ny >= len(reach[0]) {
			continue
		}

		if matrix[x][y] <= matrix[nx][ny] && !reach[nx][ny] {
			dfs(reach, matrix, nx, ny)
		}
	}
}

// @lc code=end

