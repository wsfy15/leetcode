/*
 * @lc app=leetcode.cn id=1351 lang=golang
 *
 * [1351] 统计有序矩阵中的负数
 */

// @lc code=start
func countNegatives(grid [][]int) int {
	m, res := len(grid[0]), 0
	for _, row := range grid {
		for i := m - 1; i >= 0; i-- {
			if row[i] < 0 {
				res++
			} else {
				break
			}
		}
	}

	return res
}

// @lc code=end

