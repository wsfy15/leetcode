/*
 * @lc app=leetcode.cn id=598 lang=golang
 *
 * [598] 范围求和 II
 */

// @lc code=start
func maxCount(m int, n int, ops [][]int) int {
	col, row := n, m
	for _, op := range ops {
		if op[0] < row {
			row = op[0]
		}
		if op[1] < col {
			col = op[1]
		}
	}
	return row * col
}
// @lc code=end

