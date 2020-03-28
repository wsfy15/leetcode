/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
// 也可以使用并查集
// 1. 初始时所有陆地的root指向自己
// 2. 然后遍历每块陆地，将相邻的进行合并
// 3. 最后查询有多少不同的root 这一步在第2步的时候就可以统计了

func numIslands(grid [][]byte) int {
	// 染色法，遇到陆地就+1，同时把陆地变成水
	if len(grid) == 0 {
		return 0
	}

	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				res++
				clear(grid, i, j)
			}
		}
	}
	return res
}

func clear(grid [][]byte, i, j int) {
	// dx := []int{1, -1, 0, 0}
	// dy := []int{0, 0, 1, -1}
	if grid[i][j] == '1' {
		grid[i][j] = '0'
		// for i := 0; i < 4; i++ {
		// 	x := i + dx[k]
		// 	y := j + dy[k]
		// 	if 0 <= x && x < len(grid) && 0 <= y && y < len(grid[0]) {
		// 		clear(grid, x, y)
		// 	}
		// }
		if i > 0 {
			clear(grid, i - 1, j)
		}
		if i < len(grid) - 1 {
			clear(grid, i + 1, j)
		}
		if j > 0 {
			clear(grid, i, j - 1)
		}
		if j < len(grid[0]) - 1 {
			clear(grid, i, j + 1)
		}
	}
}

func numIslands2(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	
	// DFS找到不能继续找之后，计数器加一，找剩下的1重复这个过程
	res := 0
	rows := len(grid)
	cols := len(grid[0])
	visited := make([][]bool, rows) // 记录是否访问过
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
		for j := range visited[i] {
			visited[i][j] = false
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				DFS(grid, visited, i, j)
				res++
			}
		}
	}

	return res
}

// si,sj 起点坐标
func DFS(grid [][]byte, visited [][]bool, si, sj int) {
	visited[si][sj] = true
	// 向上探索
	if si > 0 && grid[si - 1][sj] == '1' && !visited[si - 1][sj]  {
		DFS(grid, visited, si - 1, sj)
	} 

	// 向下探索
	if si < len(grid) - 1 && grid[si + 1][sj] == '1' && !visited[si + 1][sj] {
		DFS(grid, visited, si + 1, sj)
	} 

	// 向右探索
	if sj < len(grid[0]) - 1 && grid[si][sj + 1] == '1' && !visited[si][sj + 1] {
		DFS(grid, visited, si, sj + 1)
	} 

	// 向左探索
	if sj > 0 && grid[si][sj - 1] == '1' && !visited[si][sj - 1] {
		DFS(grid, visited, si, sj - 1)
	} 
}
// @lc code=end

