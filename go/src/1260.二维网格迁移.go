/*
 * @lc app=leetcode.cn id=1260 lang=golang
 *
 * [1260] 二维网格迁移
 */

// @lc code=start
func shiftGrid(grid [][]int, k int) [][]int {
	n, m := len(grid), len(grid[0])
	if n*m == k {
		return grid
	}

	nums := []int{}
	for _, row := range grid {
		nums = append(nums, row...)
	}

	k %= len(nums)
	for i, v := range nums {
		row, col := (i+k)/m, (i+k)%m
		if row >= n {
			row -= n
		}
		grid[row][col] = v
	}
	return grid
}

// @lc code=end

