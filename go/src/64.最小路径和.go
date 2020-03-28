/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum1(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])
	states := make([][]int, m) // states[i][j] 到达(i+1, j+1)的最短路径
	for i := range states {
		states[i] = make([]int, n)
	}

	sum := 0
	for i := 0; i < n; i++ { // 填充第一行
		sum += grid[0][i]
		states[0][i] = sum
	}

	sum = 0
	for i := 0; i < m; i++ { // 填充第一列
		sum += grid[i][0]
		states[i][0] = sum
	}

	for i := 1; i < m; i++ { // 逐行遍历
		for j := 1; j < n; j++ {
			min := states[i-1][j]
			if min > states[i][j-1] {
				min = states[i][j-1]
			}
			states[i][j] = min + grid[i][j]
		}
	}
	
	return states[m-1][n-1]
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])
	states := make([]int, n) // states[i] 到达(X, j+1)的最短路径 X 在后面的for循环中

	sum := 0
	for i := 0; i < n; i++ { // 初始化
		sum += grid[0][i]
		states[i] = sum
	}

	for i := 1; i < m; i++ { // 逐行遍历
		for j := 0; j < n; j++ {
			if j == 0 {
				states[0] += grid[i][0]
			} else {
				min := states[j-1] // 从左边过来
				if min > states[j] {
					min = states[j] // 从上边下来
				}
				states[j] = min + grid[i][j]
			}
		}
	}
	
	return states[n-1]
}
// @lc code=end

